package heap

import (
	"jvm_go_code/JVM/classfile"
)

type MethodRef struct {
	MemberRef
	method 		*Method
}

func newMethodRef(constantPool *ConstantPool, refInfo *classfile.CONSTANT_Methodref_info) *MethodRef {
	ref := &MethodRef{}
	ref.constantPool = constantPool
	ref.copyMemberRefInfo(&refInfo.CONSTANT_Memberref_info)
	return ref
}

// 如果方法还没有被解析, 调用 self.resolveMethodRef() 进行解析, 否则返回方法指针
func (self *MethodRef) ResolvedMethod() *Method {
	if self.method == nil {
		self.resolveMethodRef()
	}
	return self.method
}

// 分别判断 self 的归属类是不是接口, method 是否存在, method 是否能被访问, 如果都通过则返回 method
func (self *MethodRef) resolveMethodRef() {
	constantPoolClass := self.constantPool.class
	thisClass := self.ResolvedClass()
	if thisClass.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupMethod(thisClass, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(constantPoolClass) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

// 分别在类中, 接口中查找该方法
func lookupMethod(class *Class, name string, descriptor string) *Method {
	method := LookupMethodInClass(class, name, descriptor)
	if method == nil {
		method = lookupMethodInInterfaces(class.interfaces, name, descriptor)
	}
	return method
}