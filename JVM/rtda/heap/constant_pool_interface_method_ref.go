package heap

import "jvm_go_code/JVM/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method		*Method
}

func newInterfaceMethodRef(constantpool *ConstantPool, refInfo *classfile.CONSTANT_InterfaceMethodref_info) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.constantPool = constantpool
	ref.copyMemberRefInfo(&refInfo.CONSTANT_Memberref_info)
	return ref
}

// 如果方法还没有被解析, 调用 self.resolveInterfaceMethodRef() 进行解析, 否则返回方法指针
func (self *InterfaceMethodRef) ResolvedInterfaceMethod() *Method {
	if self.method == nil {
		self.resolveInterfaceMethodRef()
	}
	return self.method
}

// 分别判断 self 的归属类是不是接口, method 是否存在, method 是否能被访问, 如果都通过则返回 method
func (self *InterfaceMethodRef) resolveInterfaceMethodRef() {
	constantPoolClass := self.constantPool.class
	thisClass := self.ResolvedClass()
	if !thisClass.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	method := lookupInterfaceMethod(thisClass, self.name, self.descriptor)
	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}
	if !method.isAccessibleTo(constantPoolClass) {
		panic("java.lang.IllegalAccessError")
	}
	self.method = method
}

// 分别在接口中查找该方法
func lookupInterfaceMethod(interfaceClass *Class, name string, descriptor string) *Method {
	for _, item := range interfaceClass.methods {
		if item.name == name && item.descriptor == descriptor {
			return item
		}
	}
	return lookupMethodInInterfaces(interfaceClass.interfaces, name, descriptor)
}

