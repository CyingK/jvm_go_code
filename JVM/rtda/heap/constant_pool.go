package heap

import (
	"fmt"
	"jvm_go_code/JVM/classfile"
)

type Constant interface {}

type ConstantPool struct {
	class 	*Class			// 归属类
	consts	[]Constant		// 常量引用数组
}

// 根据下标获取常量
func (self *ConstantPool) GetConstant(index uint) Constant {
	if constant := self.consts[index]; constant != nil {
		return constant
	}
	panic(fmt.Sprintf("NoConstantFoundAtIndex: %d", index))
}

// 获取 GetConstantPool 的长度并分配空间, 创建运行时常量池, 然后循环为每个类型的常量判断类型并赋值
func newConstantPool(class *Class, classFileConstantPool classfile.ConstantPool) *ConstantPool {
	constantPoolCount := len(classFileConstantPool)
	consts := make([]Constant, constantPoolCount)
	runtimeConstantPool := &ConstantPool{
		class: 	class,
		consts: consts,
	}
	for index := 1; index < constantPoolCount; index++ {
		constantPoolInfo := classFileConstantPool[index]
		switch constantPoolInfo.(type) {
		case *classfile.CONSTANT_Integer_info:
			intInfo := constantPoolInfo.(*classfile.CONSTANT_Integer_info)
			consts[index] = intInfo.GetValue()
		case *classfile.CONSTANT_Float_info:
			floatInfo := constantPoolInfo.(*classfile.CONSTANT_Float_info)
			consts[index] = floatInfo.GetValue()
		case *classfile.CONSTANT_Long_info:
			longInfo := constantPoolInfo.(*classfile.CONSTANT_Long_info)
			consts[index] = longInfo.GetValue()
			index++
		case *classfile.CONSTANT_Double_info:
			doubleInfo := constantPoolInfo.(*classfile.CONSTANT_Double_info)
			consts[index] = doubleInfo.GetValue()
			index++
		case *classfile.CONSTANT_String_info:
			stringInfo := constantPoolInfo.(*classfile.CONSTANT_String_info)
			consts[index] = stringInfo.String()
		case *classfile.CONSTANT_Class_info:
			classInfo := constantPoolInfo.(*classfile.CONSTANT_Class_info)
			consts[index] = newClassRef(runtimeConstantPool, classInfo)
		case *classfile.CONSTANT_Fieldref_info:
			fieldRefInfo := constantPoolInfo.(*classfile.CONSTANT_Fieldref_info)
			consts[index] = newFieldRef(runtimeConstantPool, fieldRefInfo)
		case *classfile.CONSTANT_Methodref_info:
			methodRefInfo := constantPoolInfo.(*classfile.CONSTANT_Methodref_info)
			consts[index] = newMethodRef(runtimeConstantPool, methodRefInfo)
		case *classfile.CONSTANT_InterfaceMethodref_info:
			interfaceMethodRef := constantPoolInfo.(*classfile.CONSTANT_InterfaceMethodref_info)
			consts[index] = newInterfaceMethodRef(runtimeConstantPool, interfaceMethodRef)
		}
	}
	return runtimeConstantPool
}
