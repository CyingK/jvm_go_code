package heap

import (
	"jvm_go_code/array_string/classfile"
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
	initStarted				bool
}

var primitiveTypes = map[string]string{
	"void":    "V",
	"boolean": "Z",
	"byte":    "B",
	"short":   "S",
	"int":     "I",
	"long":    "J",
	"char":    "C",
	"float":   "F",
	"double":  "D",
}

// [[XXX -> [XXX
// [LXXX; -> XXX
// [I -> int
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

// [XXX -> [[XXX
// int -> [I
// XXX -> [LXXX;
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

// [XXX => [XXX
// int  => I
// XXX  => LXXX;
func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}

// [XXX  => [XXX
// LXXX; => XXX
// I     => int
func toClassName(descriptor string) string {
	if descriptor[0] == '[' {
		return descriptor
	}
	if descriptor[0] == 'L' {
		return descriptor[1: len(descriptor) - 1]
	}
	for className, _descriptor_ := range primitiveTypes {
		if _descriptor_ == descriptor {
			return className
		}
	}
	panic("Invalid descriptor: " + descriptor)
}

//--------------------------------------------------------------------构造器
// 调用 newObject(*GetClass) 创建该类的一个新对象并返回
func (self *Class) NewObject() *Object {
	return newObject(self)
}

func (self *Class) NewArray(count uint) *Object {
	if !self.IsArray() {
		panic("NotArrayClass: " + self.name)
	}
	switch self.GetName() {
	case "[Z":
		return &Object{self, make([]int8, count)}
	case "[B":
		return &Object{self, make([]int8, count)}
	case "[C":
		return &Object{self, make([]uint16, count)}
	case "[S":
		return &Object{self, make([]int16, count)}
	case "[I":
		return &Object{self, make([]int32, count)}
	case "[J":
		return &Object{self, make([]int64, count)}
	case "[F":
		return &Object{self, make([]float32, count)}
	case "[D":
		return &Object{self, make([]float64, count)}
	default:
		return &Object{self, make([]*Object, count)}
	}
}

// 创建一个新的类, 对 accessFlags, name, superClassName, interfaceNames, constantPool data, methods 进行初始化, 然后将这个类进行返回
func newClass(classFile *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = classFile.GetAccessFlags()
	class.name = classFile.GetClassName()
	class.superClassName = classFile.GetSuperClassName()
	class.interfaceNames = classFile.GetInterfaceNames()
	class.constantPool = newConstantPool(class, classFile.GetConstantPool())
	class.fields = newFields(class, classFile.GetFields())
	class.methods = newMethods(class, classFile.GetMethods())
	return class
}

//--------------------------------------------------------------------Getters
// 获取类名
func (self *Class) GetName() string {
	return self.name
}

// 获取数组类
func (self *Class) GetArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

// 获取类加载器
func (self *Class) GetClassLoader() *ClassLoader {
	return self.loader
}

// 获取 constentPool
func (self *Class) GetConstantPool() *ConstantPool {
	return self.constantPool
}

// 获取处理结果 ComponentClass
func (self *Class) GetComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}

// 获取 staticVars
func (self *Class) GetStaticVars() Slots {
	return self.staticVars
}

// 调用 self.GetStaticMethod() 获取 main() 方法
func (self *Class) GetMainMethod() *Method {
	return self.GetStaticMethod("main", "([Ljava/lang/String;)V")
}

// 获取所有方法
func (self *Class) GetAllMethods() []*Method {
	return self.methods
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

// 获取构造器
func (self *Class) GetClinitMethod() *Method {
	return self.GetStaticMethod("<clinit>", "()V")
}

// 获取父类
func (self *Class) GetSuperClass() *Class {
	return self.superClass
}

// 获取初始化标识
func (self *Class) GetInitStarted() bool {
	return self.initStarted
}

// 提取 self.name 中的包路径并返回
func (self *Class) GetPackageName() string {
	if index := strings.LastIndex(self.name, "/"); index >= 0 {
		return self.name[:index]
	}
	return ""
}

// 根据字段名和描述符查找字段
func (self *Class) getField(name string, descriptor string, isStatic bool) *Field {
	for class := self; class != nil; class = class.superClass {
		for _, item := range class.fields {
			if item.IsStatic() == isStatic &&
				item.name == name &&
				item.descriptor == descriptor {
				return item
			}
		}
	}
	return nil
}

//--------------------------------------------------------------------Setters
func (self *Class) SetInitStarted() {
	self.initStarted = true
}

//--------------------------------------------------------------------判断类方法
// 如果该类的权限是 public, 返回 true, 否则返回 false
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

// 判断当前类是不是数组类
func (self *Class) IsArray() bool {
	return self.name[0] == '['
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

// 如果 self 的权限为 public 或者 self 和 otherClass 的包路径相同, 代表 otherClass 有权限访问 self, 返回 true, 否则返回 false
func (self *Class) isAccessibleTo(otherClass *Class) bool {
	return self.IsPublic() || self.GetPackageName() == otherClass.GetPackageName()
}

// 判断一个类/接口是否可以转换成另一个类/接口
func (self *Class) isAssignableFrom(otherClass *Class) bool {
	if otherClass == self {
		return true
	}
	if !otherClass.IsArray() {
		// otherClass 不是一个数组
		if !otherClass.IsInterface() {
			// otherClass 是一个类
			if !self.IsInterface() {
				// self 也是一个类
				return otherClass.IsSubClassOf(self)
			} else {
				// self 是一个接口
				return otherClass.isImplements(self)
			}
		} else {
			// otherClass 是一个接口
			if !self.IsInterface() {
				// self 是一个类
				return self.isJlObject()
			} else {
				// self 也是一个接口
				return self.isSuperInterfaceOf(otherClass)
			}
		}
	} else {
		// otherClass 是一个数组
		if !self.IsArray() {
			// self 不是一个数组
			if !self.IsInterface() {
				// self 是一个类
				return self.isJlObject()
			} else {
				// self 是一个接口
				return self.isJlCloneable() || self.isJioSerializable()
			}
		} else {
			// self 是一个数组
			selfComponentClass := self.GetComponentClass()
			otherComponentClass := otherClass.GetComponentClass()
			return selfComponentClass == otherComponentClass ||
				otherComponentClass.isAssignableFrom(selfComponentClass)
		}
	}
}

// 判断 self 是不是 otherClass 的父类
func (self *Class) IsSuperClassOf(otherClass *Class) bool {
	return otherClass.IsSubClassOf(self)
}

// 第一重循环遍历所有父类, 第二重循环遍历类中的所有接口, 找到与 interfaceName 相同的方法则返回
func (self *Class) IsImplements(interfaceName *Class) bool {
	for class := self; class != nil; class = class.superClass {
		for _, item := range class.interfaces {
			if item == interfaceName || item.isSubInterfaceOf(interfaceName) {
				return true
			}
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

// 遍历 self 的父类, 直到没有父类为止, 每次遍历都比较被遍历对象与 otherClass, 如果两者相等则说明 self 间接继承了 otherClass, 从而返回 true, 否则返回 false
func (self *Class) IsSubClassOf(otherClass *Class) bool {
	for superClass := self.superClass; superClass != nil; superClass = superClass.superClass {
		if superClass == otherClass {
			return true
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

func (self *Class) isJlObject() bool {
	return self.name == "java/lang/Object"
}

func (self *Class) isJlCloneable() bool {
	return self.name == "java/lang/Cloneable"
}

func (self *Class) isJioSerializable() bool {
	return self.name == "java/io/Serializable"
}

func (self *Class) isSuperInterfaceOf(_interface_ *Class) bool {
	return _interface_.isSubInterfaceOf(self)
}