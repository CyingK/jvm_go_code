package heap

import (
	"jvm_go_code/JVM/classfile"
)

type Method struct {
	ClassMember
	maxStack                uint
	maxLocals               uint
	code                    []byte
	exceptionTable          ExceptionTable
	lineNumberTable         *classfile.ATTRIBUTE_LINE_NUMBER_TABLE
	exceptions              *classfile.ATTRIBUTE_EXCEPTIONS
	parameterAnnotationData []byte
	annotationDefaultData	[]byte
	parsedDescriptor        *MethodDescriptor
	argSlotCount            uint
}

//--------------------------------------------------------------------构造器

func newMethods(class *Class, classFileMethods []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(classFileMethods))
	for index, item := range classFileMethods {
		methods[index] = newMethod(class, item)
	}
	return methods
}

func newMethod(class *Class, classFileMethod *classfile.MemberInfo) *Method {
	method := &Method{}
	method.class = class
	method.copyMemberInfo(classFileMethod)
	method.copyAttributes(classFileMethod)
	methodDescriptor := parseMethodDescriptor(method.descriptor)
	method.parsedDescriptor = methodDescriptor
	method.calcArgSlotCount(methodDescriptor.parameterTypes)
	if method.IsNative() {
		method.injectCodeAttributes(methodDescriptor.returnType)
	}
	return method
}

//--------------------------------------------------------------------Getters

// 获取 maxLocals
func (self *Method) GetMaxLocals() uint {
	return self.maxLocals
}

func (self *Method) GetParsedDescriptor() *MethodDescriptor {
	return self.parsedDescriptor
}

// 获取 maxStack
func (self *Method) GetMaxStack() uint {
	return self.maxStack
}

// 获取 code
func (self *Method) GetCode() []byte {
	return self.code
}

// 获取 argSlotCount
func (self *Method) GetArgSlotCount() uint {
	return self.argSlotCount
}

//--------------------------------------------------------------------判断类方法

func (self *Method) IsSynchronized() bool {
	return 0 != self.accessFlags&ACC_SYNCHRONIZED
}

func (self *Method) IsBridge() bool {
	return 0 != self.accessFlags&ACC_BRIDGE
}

func (self *Method) IsVarargs() bool {
	return 0 != self.accessFlags&ACC_VARARGS
}

func (self *Method) IsNative() bool {
	return 0 != self.accessFlags&ACC_NATIVE
}

func (self *Method) IsAbstract() bool {
	return 0 != self.accessFlags&ACC_ABSTRACT
}

func (self *Method) IsStrict() bool {
	return 0 != self.accessFlags&ACC_STRICT
}

//--------------------------------------------------------------------功能类方法

// 计算局部变量表所需插槽数, 对每个参数都计一个插槽, Long 和 Double 类型要多占一个, 其次如果是非静态方法还要留一个给 this
func (self *Method) calcArgSlotCount(paramTypes []string) {
	for _, item := range paramTypes {
		self.argSlotCount++
		if item == "J" || item == "D" {
			self.argSlotCount++
		}
	}
	if !self.IsStatic() {
		self.argSlotCount++
	}
}

// 复制属性
func (self *Method) copyAttributes(classFileMethod *classfile.MemberInfo) {
	if codeAttr := classFileMethod.GetCodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.GetMaxStack()
		self.maxLocals = codeAttr.GetMaxLocals()
		self.code = codeAttr.GetCode()
		self.lineNumberTable = codeAttr.GetLineNumberAttribute()
		self.exceptionTable = newExceptionTable(codeAttr.GetExceptionTable(), self.GetClass().GetConstantPool())
	}
	self.exceptions = classFileMethod.GetAttributeExceptions()
	self.annotationData = classFileMethod.GetRuntimeVisibleAnnotationsAttributeData()
	self.parameterAnnotationData = classFileMethod.GetRuntimeVisibleParameterAnnotationsAttributeData()
	self.annotationDefaultData = classFileMethod.GetAnnotationDefaultAttributeData()
}

func (self *Method) injectCodeAttributes(returnType string) {
	self.maxStack = 4
	self.maxLocals = self.argSlotCount
	switch returnType[0] {
	case 'V':
		self.code = []byte{0xFE, 0xB1}
	case 'L', '[':
		self.code = []byte{0xFE, 0xB0}
	case 'D':
		self.code = []byte{0xFE, 0xAF}
	case 'F':
		self.code = []byte{0xFE, 0xAE}
	case 'J':
		self.code = []byte{0xFE, 0xAD}
	default:
		self.code = []byte{0xFE, 0xAC}
	}
}

func (self *Method) FindExceptionHandler(exClass *Class, pc int) int {
	handler := self.exceptionTable.findExceptionHandler(exClass, pc)
	if handler != nil {
		return handler.handlerPc
	}
	return -1
}

func (self *Method) GetLineNumber(pc int) int {
	if self.IsNative() {
		return -2
	}
	if self.lineNumberTable == nil {
		return -1
	}
	return self.lineNumberTable.GetLineNumber(pc)
}

func (self *Method) isConstructor() bool {
	return !self.IsStatic() && self.name == "<init>"
}

func (self *Method) isClinit() bool {
	return self.IsStatic() && self.name == "<clinit>"
}

func (self *Method) GetParameterTypes() []*Class {
	if self.argSlotCount == 0 {
		return nil
	}
	paramTypes := self.parsedDescriptor.parameterTypes
	paramClasses := make([]*Class, len(paramTypes))
	for i, paramType := range paramTypes {
		paramClassName := toClassName(paramType)
		paramClasses[i] = self.class.loader.LoadClass(paramClassName)
	}
	return paramClasses
}

func (self *Method) GetAnnotationDefaultData() []byte {
	return self.annotationDefaultData
}

func (self *Method) GetExceptionTypes() []*Class {
	if self.exceptions == nil {
		return nil
	}
	exceptionIndexTable := self.exceptions.ExceptionIndexTable()
	exceptionClasses := make([]*Class, len(exceptionIndexTable))
	constantPool := self.class.constantPool
	for index, item := range exceptionIndexTable {
		classRef := constantPool.GetConstant(uint(item)).(*ClassRef)
		exceptionClasses[index] = classRef.ResolvedClass()
	}
	return exceptionClasses
}

func (self *Method) GetReturnType() *Class {
	returnType := self.parsedDescriptor.returnType
	returnClassName := toClassName(returnType)
	return self.class.loader.LoadClass(returnClassName)
}

func (self *Method) GetAccessFlags() uint16 {
	return self.accessFlags
}


func (self *Method) GetAnnotationData() []byte {
	return self.annotationData
}

func (self *Method) GetParameterAnnotationData() []byte {
	return self.parameterAnnotationData
}