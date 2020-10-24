package heap

import (
	"jvm_go_code/JVM/classfile"
	"strings"
)

type Class struct {
	accessFlags				classfile.U2		// 访问标识
	name					string				// 该类的全限定类名
	superClassName			string				// 父类的全限定类名
	interfaceNames			[]string			// 接口的全限定类名数组
	constantPool			*ConstantPool		// 常量池
	fields					[]*Field			// 属性引用数组
	methods					[]*Method			// 方法引用数组
	sourceFile				string				// 源文件名
	loader					*ClassLoader		// 类加载器
	superClass				*Class				// 父类引用
	interfaces				[]*Class			// 接口引用数组
	instanceSlotCount		uint				// 接口数
	staticSlotCount			uint				// 静态属性数
	staticVars				Slots				// 静态属性引用数组
	initStarted				bool				// 已经初始化标识
	java_class				*Object				// 指向实例的引用
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

// 获取该类的访问标识
func (self *Class) GetAccessFlags() classfile.U2 {
	return self.accessFlags
}

// 获取该类的所有方法
func (self *Class) GetAllMethods() []*Method {
	return self.methods
}

// 获取该类的一个数组类, 即在该类基础上套了一层数组
func (self *Class) GetArrayClass() *Class {
	arrayClassName := getArrayClassName(self.name)
	return self.loader.LoadClass(arrayClassName)
}

// 获取该类的类加载器
func (self *Class) GetClassLoader() *ClassLoader {
	return self.loader
}

// 获取该类的 <clinit> 方法
func (self *Class) GetClinitMethod() *Method {
	return self.getMethod("<clinit>", "()V", true)
}

// 获取该数组类的原本类型
func (self *Class) GetComponentClass() *Class {
	componentClassName := getComponentClassName(self.name)
	return self.loader.LoadClass(componentClassName)
}

// 获取该类的运行时常量池
func (self *Class) GetConstantPool() *ConstantPool {
	return self.constantPool
}

// 获取该类的构造器
func (self *Class) GetConstructor(descriptor string) *Method {
	return self.GetInstanceMethod("<init>", descriptor)
}

// 获取该类的所有构造器(根据 publicOnly 判断是否要获取 private 的构造器)
func (self *Class) GetConstructors(publicOnly bool) []*Method {
	methods := make([]*Method, 0, len(self.methods))
	for _, method := range self.methods {
		if !method.isClinit() && !method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				methods = append(methods, method)
			}
		}
	}
	return methods
}

// 获取该类的所有字段(根据 publicOnly 判断是否要获取 private 的字段)
func (self *Class) GetFields(publicOnly bool) []*Field {
	if publicOnly {
		publicFields := make([]*Field, 0, len(self.fields))
		for _, field := range self.fields {
			if field.IsPublic() {
				publicFields = append(publicFields, field)
			}
		}
		return publicFields
	} else {
		return self.fields
	}
}

// 获取该类的初始化标识
func (self *Class) GetInitStarted() bool {
	return self.initStarted
}

// 根据方法名, 描述信息获取该类的实例方法
func (self *Class) GetInstanceMethod(name string, descriptor string) *Method {
	return self.getMethod(name, descriptor, false)
}

// 获取该类实现的所有接口
func (self *Class) GetInterfaces() []*Class {
	return self.interfaces
}

// 获取该类的 java_class
func (self *Class) GetJavaClass() *Object {
	return self.java_class
}

// 获取该类的全类名(用 . 替换 /)
func (self *Class) GetJavaName() string {
	return strings.Replace(self.name, "/", ".", -1)
}

// 获取该该类的 main() 方法
func (self *Class) GetMainMethod() *Method {
	return self.getMethod("main", "([Ljava/lang/String;)V", true)
}

// 获取该类的所有方法(根据publicOnly)确定是否要获取私有方法
func (self *Class) GetMethods(publicOnly bool) []*Method {
	methods := make([]*Method, 0, len(self.methods))
	for _, method := range self.methods {
		if !method.isClinit() && !method.isConstructor() {
			if !publicOnly || method.IsPublic() {
				methods = append(methods, method)
			}
		}
	}
	return methods
}

// 获取该类的类名
func (self *Class) GetName() string {
	return self.name
}

// 获取该类的包名
func (self *Class) GetPackageName() string {
	if index := strings.LastIndex(self.name, "/"); index >= 0 {
		return self.name[:index]
	}
	return ""
}

// 获取该类的源文件名
func (self *Class) GetSourceFile() string {
	return self.sourceFile
}

// 获取该类中指定名称, 描述符的静态方法: 遍历 self 中的每一个方法, 如果是被标记为静态, 且方法名和方法描述与实参一致, 则返回
func (self *Class) GetStaticMethod(name string, descriptor string) *Method {
	for _, item := range self.methods {
		if item.IsStatic() && item.name == name && item.descriptor == descriptor {
			return item
		}
	}
	return nil
}

// 获取该类指定名称, 描述信息的静态字段
func (self *Class) GetStaticRefVar(fieldName string, fieldDescriptor string) *Object {
	field := self.getField(fieldName, fieldDescriptor, true)
	return self.staticVars.GetRef(field.slotId)
}

// 获取所有静态引用
func (self *Class) GetStaticVars() Slots {
	return self.staticVars
}

// 获取父类
func (self *Class) GetSuperClass() *Class {
	return self.superClass
}

// 如果该类是一个抽象类, 返回 true, 否则返回 false
func (self *Class) IsAbstract() bool {
	return 0 != self.accessFlags & ACC_ABSTRACT
}

// 如果该类是一个注解, 返回 true, 否则返回 false
func (self *Class) IsAnnotation() bool {
	return 0 != self.accessFlags & ACC_ANNOTATION
}

// 判断当前类是不是数组类
func (self *Class) IsArrayClass() bool {
	return self.name[0] == '['
}

// self 可以转换成 otherClass: true, 反之 false
func (self *Class) IsAssignableFrom(otherClass *Class) bool {
	if self == otherClass {
		return true
	}
	if !otherClass.IsArrayClass() {
		if !otherClass.IsInterface() {
			if !self.IsInterface() {
				return otherClass.IsSubClassOf(self)
			} else {
				return otherClass.IsImplements(self)
			}
		} else {
			if !self.IsInterface() {
				return self.isJavaLangObject()
			} else {
				return self.isSuperInterfaceOf(otherClass)
			}
		}
	} else {
		if !self.IsArrayClass() {
			if !self.IsInterface() {
				return self.isJavaLangObject()
			} else {
				return self.isJavaLangCloneable() || self.isJavaIoSerializable()
			}
		} else {
			sc := otherClass.GetComponentClass()
			tc := self.GetComponentClass()
			return sc == tc || tc.IsAssignableFrom(sc)
		}
	}
}

// 如果该类是一个枚举类, 返回 true, 否则返回 false
func (self *Class) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}

// 如果该类被标记为 final, 返回 true, 否则返回 false
func (self *Class) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
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

// 如果该类是一个 interface, 返回 true, 否则返回 false
func (self *Class) IsInterface() bool {
	return 0 != self.accessFlags & ACC_INTERFACE
}

// 如果该类是一个基本数据类型类, 返回 true
func (self *Class) IsPrimitive() bool {
	_, ok := primitiveTypes[self.name]
	return ok
}

// 如果该类的权限是 public, 返回 true, 否则返回 false
func (self *Class) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
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

// 如果该类被标记为 super, 返回 true, 否则返回 false
func (self *Class) IsSuper() bool {
	return 0 != self.accessFlags & ACC_SUPER
}

// 判断 self 是不是 otherClass 的父类
func (self *Class) IsSuperClassOf(otherClass *Class) bool {
	return otherClass.IsSubClassOf(self)
}

// 如果该类被标记为 synthetic, 返回 true, 否则返回 false
func (self *Class) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

// 创建该类的一个数组, 共 count 个元素, 最后返回
func (self *Class) NewArray(count uint) *Object {
	if !self.IsArrayClass() {
		panic("NotArrayClass: " + self.name)
	}
	switch self.GetName() {
	case "[Z":
		return &Object{self, make([]int8, count), nil}
	case "[B":
		return &Object{self, make([]int8, count), nil}
	case "[C":
		return &Object{self, make([]uint16, count), nil}
	case "[S":
		return &Object{self, make([]int16, count), nil}
	case "[I":
		return &Object{self, make([]int32, count), nil}
	case "[J":
		return &Object{self, make([]int64, count), nil}
	case "[F":
		return &Object{self, make([]float32, count), nil}
	case "[D":
		return &Object{self, make([]float64, count), nil}
	default:
		return &Object{self, make([]*Object, count), nil}
	}
}

// 调用 newObject(*GetClass) 创建该类的一个新对象并返回
func (self *Class) NewObject() *Object {
	return newObject(self)
}

// 将该类是否初始化的标志置 true
func (self *Class) SetInitStarted() {
	self.initStarted = true
}

// 根据字段名, 字段描述, 获取该类中的一个字段, 并将 ref 引用赋值给字段
func (self *Class) SetStaticRefVar(fieldName, fieldDescriptor string, ref *Object) {
	field := self.getField(fieldName, fieldDescriptor, true)
	self.staticVars.SetRef(field.slotId, ref)
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

// 根据方法名, 方法描述, 是否静态从该类中获取方法并返回
func (self *Class) getMethod(name, descriptor string, isStatic bool) *Method {
	for c := self; c != nil; c = c.superClass {
		for _, method := range c.methods {
			if method.IsStatic() == isStatic &&
				method.name == name &&
				method.descriptor == descriptor {

				return method
			}
		}
	}
	return nil
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
	if !otherClass.IsArrayClass() {
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
				return self.isJavaLangObject()
			} else {
				// self 也是一个接口
				return self.isSuperInterfaceOf(otherClass)
			}
		}
	} else {
		// otherClass 是一个数组
		if !self.IsArrayClass() {
			// self 不是一个数组
			if !self.IsInterface() {
				// self 是一个类
				return self.isJavaLangObject()
			} else {
				// self 是一个接口
				return self.isJavaLangCloneable() || self.isJavaIoSerializable()
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

// 该类是 _interface_ 的实现类: true, 反之: false
// 第一重循环遍历 self 及其所有父类, 得到 thisClass. 第二重循环遍历 thisClass 中的每一个接口. 如果有一个接口与 interfaceName 相同,
// 或者是 interfaceName 的子接口，则说明 self 实现了 interfaceName，返回 true, 否则返回 false
func (self *Class) isImplements(_interface_ *Class) bool {
	for thisClass := self; thisClass != nil; thisClass = thisClass.superClass {
		for _, item := range thisClass.interfaces {
			if item == _interface_ || item.isSubInterfaceOf(_interface_) {
				return true
			}
		}
	}
	return false
}

// 当前类为 java/io/Serializable.class: true, 反之: false
func (self *Class) isJavaIoSerializable() bool {
	return self.name == "java/io/Serializable"
}

// 当前类为 java/lang/Cloneable.class: true, 反之: false
func (self *Class) isJavaLangCloneable() bool {
	return self.name == "java/lang/Cloneable"
}

// 当前类为 java/lang/Object.class: true, 反之: false
func (self *Class) isJavaLangObject() bool {
	return self.name == "java/lang/Object"
}

// self 是 _interface_ 的子接口: true, 反之: false
func (self *Class) isSubInterfaceOf(_interface_ *Class) bool {
	// 遍历 self 的所有父接口, 找到 _interface_ 匹配的则返回
	for _, superInterface := range self.interfaces {
		if superInterface == _interface_ || superInterface.isSubInterfaceOf(_interface_) {
			return true
		}
	}
	return false
}

// self 是 _interface_ 的父接口: true, 反之: false
func (self *Class) isSuperInterfaceOf(_interface_ *Class) bool {
	return _interface_.isSubInterfaceOf(self)
}

// 创建 byte 数组
func NewByteArray(loader *ClassLoader, bytes []int8) *Object {
	return &Object{loader.LoadClass("[B"), bytes, nil}
}

// 在原该类型的基础上加上 [, 使其成为一个数组类
func getArrayClassName(className string) string {
	return "[" + toDescriptor(className)
}

// 在数组类的基础上剥去一层数组
func getComponentClassName(className string) string {
	if className[0] == '[' {
		componentTypeDescriptor := className[1:]
		return toClassName(componentTypeDescriptor)
	}
	panic("Not array: " + className)
}

// 获取该类源文件名
func getSourceFile(classFile *classfile.ClassFile) string {
	if sourceFileAttribute := classFile.GetSourceFileAttribute(); sourceFileAttribute != nil {
		return sourceFileAttribute.GetFileName()
	}
	return "Unkown"
}

// 创建一个新的类, 进行初始化, 然后将这个类进行返回
func newClass(classFile *classfile.ClassFile) *Class {
	class := &Class{}
	class.accessFlags = classFile.GetAccessFlags()
	class.name = classFile.GetClassName()
	class.superClassName = classFile.GetSuperClassName()
	class.interfaceNames = classFile.GetInterfaceNames()
	class.constantPool = newConstantPool(class, classFile.GetConstantPool())
	class.fields = newFields(class, classFile.GetFields())
	class.methods = newMethods(class, classFile.GetMethods())
	class.sourceFile = getSourceFile(classFile)
	return class
}

// 将 classes 转为 Class 数组
func toClassArray(classLoader *ClassLoader, classes []*Class) *Object {
	arrLen := len(classes)
	classArrClass := classLoader.LoadClass("java/lang/Class").GetArrayClass()
	classArr := classArrClass.NewArray(uint(arrLen))
	if arrLen > 0 {
		classObjs := classArr.GetRefs()
		for i, class := range classes {
			classObjs[i] = class.GetJavaClass()
		}
	}
	return classArr
}

// 将数组类直接返回, 基本数据类还原成 Java 中所使用的关键字, 引用类型去除全限定名并返回
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

// 将数组类直接返回, 基本数据类转换成对应的大写标识, 引用类型改为全限定名并返回
func toDescriptor(className string) string {
	if className[0] == '[' {
		return className
	}
	if d, ok := primitiveTypes[className]; ok {
		return d
	}
	return "L" + className + ";"
}
