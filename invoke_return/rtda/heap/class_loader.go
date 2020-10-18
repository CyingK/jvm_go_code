package heap

import (
	"fmt"
	"jvm_go_code/invoke_return/classfile"
	"jvm_go_code/invoke_return/classpath"
)

type ClassLoader struct {
	classPath			*classpath.ClassPath
	classMap			map[string]*Class
}

// 创建一个新的 ClassLoader 并返回
func NewClassLoader(classPath *classpath.ClassPath) *ClassLoader {
	return &ClassLoader{
		classPath: 		classPath,
		classMap: 		make(map[string]*Class),
	}
}

// 载入一个类, 首先去 classMap 找到 className 对应的类, 如果找到则直接返回, 否则调用 self.loadNonArrayClass(string) 来加载
func (self *ClassLoader) LoadClass(className string) *Class {
	if class, ok := self.classMap[className]; ok {
		return class
	}
	return self.loadNonArrayClass(className)
}

// 载入一个类, 首先读取并编译 class, 然后对其进行验证, 初始化, 最后返回
func (self *ClassLoader) loadNonArrayClass(name string) *Class {
	data, entry := self.readClass(name)
	class := self.defineClass(data)
	link(class)
	return class
}

// 从 classPath 读入类数据返回
func (self *ClassLoader) readClass(name string) ([]byte, classpath.Entry) {
	data, entry, err := self.classPath.ReadClass(name)
	if err != nil {
		panic("java.lang.ClassNotFoundException: " + name)
	}
	return data, entry
}

// 编译 class, 指定其 loader 为 self, 装载其父类, 接口, 完成后将映射添加到 classMap 中, 最后返回此类
func (self *ClassLoader) defineClass(data []byte) *Class {
	class := parseClass(data)
	class.loader = self
	resolveSuperClass(class)
	resolveInterfaces(class)
	self.classMap[class.name] = class
	return class
}

// 装载接口, 根据 interfaceNames 的长度分配内存空间, 然后循环载入并装入 self.interfaces
func resolveInterfaces(class *Class) {
	interfaceCount := len(class.interfaceNames)
	if interfaceCount > 0 {
		class.interfaces = make([]*Class, interfaceCount)
		for index, item := range class.interfaceNames {
			class.interfaces[index] = class.loader.LoadClass(item)
		}
	}
}

// 装载父类, 判断该类是不是 Object 类, 如果不是则载入 class.superClassName, 并装入 self.superClass
func resolveSuperClass(class *Class) {
	if class.name != "java/lang/Object" {
		class.superClass = class.loader.LoadClass(class.superClassName)
	}
}

// 编译 class 的二进制数据, 调用 newClass() 并返回
func parseClass(data []byte) *Class {
	classFile, err := classfile.Parse(data)
	if err != nil {
		panic("java.lang.ClassFormatError")
	}
	return newClass(classFile)
}

// 链接过程, 分为检查和初始化两个环节
func link(class *Class) {
	verify(class)
	prepare(class)
}

// 初始化环节分为初始化属性
func prepare(class *Class) {
	calcInstanceFieldSlotIds(class)
	calcStaticFieldSlotIds(class)
	allocAndInitStaticVars(class)
}

// 为所有静态属性分配空间并初始化
func allocAndInitStaticVars(class *Class) {
	class.staticVars = newSlots(class.staticSlotCount)
	for _, item := range class.fields {
		if item.IsStatic() && item.IsFinal() {
			initStaticFinalVar(class, item)
		}
	}
}

// 初始化静态属性
func initStaticFinalVar(class *Class, field *Field) {
	vars := class.staticVars
	constantPool := class.constantPool
	constantPoolIndex := field.ConstValueIndex()
	slotId := field.SlotId()
	if constantPoolIndex > 0 {
		switch field.Descriptor() {
		case "Z", "B", "C", "S", "I":
			value := constantPool.GetConstant(constantPoolIndex).(int32)
			vars.SetInt(slotId, value)
		case "J":
			value := constantPool.GetConstant(constantPoolIndex).(int64)
			vars.SetLong(slotId, value)
		case "F":
			value := constantPool.GetConstant(constantPoolIndex).(float32)
			vars.SetFloat(slotId, value)
		case "D":
			value := constantPool.GetConstant(constantPoolIndex).(float64)
			vars.SetDouble(slotId, value)
		case "Ljava/lang/String;":
			panic("class_loader.initStaticFinalVar()")
		}
	}
}

// 首先如果有父类, 则 slotId 从父类的 inistanceSlotCount 开始计算, 否则从零开始. 然后对 self 中的所有静态属性进行计算/编号
func calcStaticFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, item := range class.fields {
		if item.IsStatic() {
			item.slotId = slotId
			slotId++
			if item.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.staticSlotCount = slotId
}

// 首先如果有父类, 则 slotId 从父类的 inistanceSlotCount 开始计算, 否则从零开始. 然后对 self 中的所有非静态属性进行计算/编号
func calcInstanceFieldSlotIds(class *Class) {
	slotId := uint(0)
	if class.superClass != nil {
		slotId = class.superClass.instanceSlotCount
	}
	for _, item := range class.fields {
		if !item.IsStatic() {
			item.slotId = slotId
			slotId++
			if item.isLongOrDouble() {
				slotId++
			}
		}
	}
	class.instanceSlotCount = slotId
}

func verify(class *Class) {

}
