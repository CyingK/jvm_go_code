package heap

import (
	"fmt"
	"jvm_go_code/array_string/classfile"
)

type Method struct {
	ClassMember
	maxStack		uint
	maxLocals		uint
	code			[]byte
	argSlotCount	uint
}

//--------------------------------------------------------------------构造器

func newMethods(class *Class, classFileMethod []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(classFileMethod))
	for index, item := range classFileMethod {
		methods[index] = &Method{}
		methods[index].class = class
		methods[index].copyMemberInfo(item)
		methods[index].copyAttributes(item)
		methods[index].calcArgSlotCount()
	}
	return methods
}

//--------------------------------------------------------------------Getters

// 获取 maxLocals
func (self *Method) GetMaxLocals() uint {
	return self.maxLocals
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
func (self *Method) calcArgSlotCount() {
	fmt.Printf("初始化方法... 类名：%v 方法名：%v 方法描述：%v\n", self.GetClass().GetName(), self.GetName(), self.Descriptor())
	parsedDescriptor := parseMethodDescriptor(self.descriptor)
	for _, item := range parsedDescriptor.parameterTypes {
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
	}
}