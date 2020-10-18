package heap

import "jvm_go_code/invoke_return/classfile"

type Method struct {
	ClassMember
	maxStack		uint
	maxLocals		uint
	code			[]byte
}

func newMethods(class *Class, classFileMethod []*classfile.MemberInfo) []*Method {
	methods := make([]*Method, len(classFileMethod))
	for index, item := range classFileMethod {
		methods[index] = &Method{}
		methods[index].class = class
		methods[index].copyMemberInfo(item)
		methods[index].copyAttributes(item)
	}
	return methods
}

func (self *Method) copyAttributes(classFileMethod *classfile.MemberInfo) {
	if codeAttr := classFileMethod.CodeAttribute(); codeAttr != nil {
		self.maxStack = codeAttr.MaxStack()
		self.maxLocals = codeAttr.MaxLocals()
		self.code = codeAttr.Code()
	}
}

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

func (self *Method) MaxLocals() uint {
	return self.maxLocals
}

func (self *Method) MaxStack() uint {
	return self.maxStack
}

func (self *Method) Code() []byte {
	return self.code
}