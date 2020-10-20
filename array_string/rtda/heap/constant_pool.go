package heap

import (
	"fmt"
	"jvm_go_code/array_string/classfile"
)

type Constant interface {

}

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
		case *classfile.CONSTANT_INTEGER_INFO:
			intInfo := constantPoolInfo.(*classfile.CONSTANT_INTEGER_INFO)
			consts[index] = intInfo.GetValue()
		case *classfile.CONSTANT_FLOAT_INFO:
			floatInfo := constantPoolInfo.(*classfile.CONSTANT_FLOAT_INFO)
			consts[index] = floatInfo.GetValue()
		case *classfile.CONSTANT_LONG_INFO:
			longInfo := constantPoolInfo.(*classfile.CONSTANT_LONG_INFO)
			consts[index] = longInfo.GetValue()
			index++
		case *classfile.CONSTANT_DOUBLE_INFO:
			doubleInfo := constantPoolInfo.(*classfile.CONSTANT_DOUBLE_INFO)
			consts[index] = doubleInfo.GetValue()
			index++
		case *classfile.CONSTANT_STRING_INFO:
			stringInfo := constantPoolInfo.(*classfile.CONSTANT_STRING_INFO)
			consts[index] = stringInfo.String()
		case *classfile.CONSTANT_CLASS_INFO:
			classInfo := constantPoolInfo.(*classfile.CONSTANT_CLASS_INFO)
			consts[index] = newClassRef(runtimeConstantPool, classInfo)
		case *classfile.CONSTANT_FIELD_REF_INFO:
			fieldRefInfo := constantPoolInfo.(*classfile.CONSTANT_FIELD_REF_INFO)
			consts[index] = newFieldRef(runtimeConstantPool, fieldRefInfo)
		case *classfile.CONSTANT_METHOD_REF_INFO:
			methodRefInfo := constantPoolInfo.(*classfile.CONSTANT_METHOD_REF_INFO)
			consts[index] = newMethodRef(runtimeConstantPool, methodRefInfo)
		case *classfile.CONSTANT_INTERFACE_METHOD_REF_INFO:
			interfaceMethodRef := constantPoolInfo.(*classfile.CONSTANT_INTERFACE_METHOD_REF_INFO)
			consts[index] = newInterfaceMethodRef(runtimeConstantPool, interfaceMethodRef)
		}
	}
	return runtimeConstantPool
}
