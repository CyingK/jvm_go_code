package heap

import (
	"jvm_go_code/invoke_return/classfile"
	"strings"
)

type Class struct {
	accessFlags				uint16				// 访问标识
	name					string				// 本类的全限定类名
	superClassName			string				// 父类的全限定类名
	interfaceNames			[]string			// 接口的全限定类名数组
	constantPool			*ConstantPool		// 常量池
	fields					[]*Field			// 属性引用数组
	methods					[]*Method			// 方法引用数组
	loader					*ClassLoader		// 类加载器
	superClass				*Class				// 父类引用
	interfaces				[]*Class			// 接口引用数组
	instanceSlotCount		uint				// 接口数
	staticSlotCount			uint				// 静态属性数
	staticVars				Slots				// 静态属性引用数组
}

// 创建一个新的类, 对 accessFlags, name, superClassName, interfaceNames, constantPool fields, methods 进行初始化, 然后将这个类进行返回
func newClass(classFile *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = classFile.AccessFlags()
	class.name = classFile.ClassName()
	class.superClassName = classFile.SuperClassName()
	class.interfaceNames = classFile.InterfaceNames()
	class.constantPool = newConstantPool(class, classFile.ConstantPool())
	class.fields = newFields(class, classFile.Fields())
	class.methods = newMethods(class, classFile.Methods())
	return class
}

func (self *Class) GetName() string {
	return self.name
}

// 调用 newObject(*Class) 创建该类的一个新对象并返回
func (self *Class) NewObject() *Object {
	return newObject(self)
}

// 如果该类的权限是 public, 返回 true, 否则返回 false
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

// 如果该类被标记为 final, 返回 true, 否则返回 false
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

// 如果该类被标记为 super, 返回 true, 否则返回 false
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags & ACC_SUPER
}

// 如果该类是一个 interface, 返回 true, 否则返回 false
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags & ACC_INTERFACE
}

// 如果该类是一个抽象类, 返回 true, 否则返回 false
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

// 如果该类被标记为 synthetic, 返回 true, 否则返回 false
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

// 如果该类是一个注解, 返回 true, 否则返回 false
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags & ACC_ANNOTATION
}

// 如果该类是一个枚举类, 返回 true, 否则返回 false
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}

// 获取 constentPool
func (self *Class) ConstantPool() *ConstantPool {
	return self.constantPool
}

// 获取 staticVars
func (self *Class) StaticVars() Slots {
	return self.staticVars
}

// 如果 self 的权限为 public 或者 self 和 otherClass 的包路径相同, 代表 otherClass 有权限访问 self, 返回 true, 否则返回 false
func (self *Class) isAccessibleTo(otherClass *Class) bool {
	return self.IsPublic() || self.getPackageName() == otherClass.getPackageName()
}

// 提取 self.name 中的包路径并返回
func (self *Class) getPackageName() string {
	if index := strings.LastIndex(self.name, "/"); index >= 0 {
		return self.name[:index]
	}
	return ""
}

// 调用 self.GetStaticMethod() 获取 main() 方法
func (self *Class) GetMainMethod() *Method {
	return self.GetStaticMethod("main", "([Ljava/lang/String;)V")
}

//遍历 self 中的每一个方法, 如果是被标记为静态, 且方法名和方法描述与实参一致, 则返回
func (self *Class) GetStaticMethod(name string, descriptor string) *Method {
	for _, item := range self.methods {
		if item.IsStatic() && item.name == name && item.descriptor == descriptor {
			return item
		}
	}
	return nil
}

// 判断一个类/接口是否可以转换成另一个类/接口
func (self *Class) isAssignableFrom(otherClass *Class) bool {
	if self == otherClass {
		return true
	}
	if !self.IsInterface() {
	 	return otherClass.isSubClassOf(self)
	} else {
		return otherClass.isImplements(self)
	}
}

// 遍历 self 的父类, 直到没有父类为止, 每次遍历都比较被遍历对象与 otherClass, 如果两者相等则说明 self 间接继承了 otherClass, 从而返回 true, 否则返回 false
func (self *Class) isSubClassOf(otherClass *Class) bool {
	for superClass := self.superClass; superClass != nil; superClass = superClass.superClass {
		if superClass == otherClass {
			return true
		}
	}
	return false
}
// 第一重循环遍历 self 及其所有父类, 得到 thisClass. 第二重循环遍历 thisClass 中的每一个接口. 如果有一个接口与 interfaceName 相同，或者是 interfaceName 的子接口，则说明 self 实现了 interfaceName，返回 true, 否则返回 false
func (self *Class) isImplements(interfaceName *Class) bool {
	for thisClass := self; thisClass != nil; thisClass = thisClass.superClass {
		for _, item := range thisClass.interfaces {
			if item == interfaceName || item.isSubInterfaceOf(interfaceName) {
				return true
			}
		}
	}
	return false
}

// 利用遍历 + 递归的方法，将 self 与父接口中继承的所有接口与 interfaceName 做对比
func (self *Class) isSubInterfaceOf(interfaceName *Class) bool {
	for _, superInterface := range self.interfaces {
		if superInterface == interfaceName || superInterface.isSubInterfaceOf(interfaceName) {
			return true
		}
	}
	return false
}

