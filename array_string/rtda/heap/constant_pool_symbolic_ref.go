package heap

type SymRef struct {
	constantPool			*ConstantPool	// 运行时常量池
	className				string			// 全限定类名
	class					*Class			// 类
}

// 如果 self.class 为空, 则调用 self.resolveClassRef() 装载 self.className, 然后返回
func (self *SymRef) ResolvedClass() *Class {
	if self.class == nil {
		self.resolveClassRef()
	}
	return self.class
}

func (self *SymRef) resolveClassRef() {
	constantPoolClass := self.constantPool.class
	class := constantPoolClass.loader.LoadClass(self.className)
	if !class.isAccessibleTo(constantPoolClass) {
		panic("java.lang.IllegalAccessError")
	}
	self.class = class
}

func (self *SymRef) GetClassName() string {
	return self.className
}