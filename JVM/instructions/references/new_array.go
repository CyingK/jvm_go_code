package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
	"strconv"
)

const (
	AT_BOOLEAN	= 4
	AT_CHAR	= 5
	AT_FLOAT	= 6
	AT_DOUBLE	= 7
	AT_BYTE	= 8
	AT_SHORT	= 9
	AT_INT		= 10
	AT_LONG	= 11
)

// 创建基本类型数组，包括boolean[]、byte[]、char[]、short[]、int[]、long[]、float[]和double[]
type NEW_ARRAY struct {
	atype uint8
}

func (self *NEW_ARRAY) GetOperands(reader *base.ByteCodeReader) {
	self.atype = reader.ReadUint8()
}

func (self *NEW_ARRAY) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	classLoader := frame.GetMethod().GetClass().GetClassLoader()
	arrClass := getPrimitiveArrayClass(classLoader, self.atype)
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
}

func (self *NEW_ARRAY) String() string {
	return "{type：new_array; type: " + strconv.Itoa(int(self.atype)) + "}\t"
}

func getPrimitiveArrayClass(loader *heap.ClassLoader, atype uint8) *heap.Class {
	switch atype {
	case AT_BOOLEAN:
		return loader.LoadClass("[Z")
	case AT_BYTE:
		return loader.LoadClass("[B")
	case AT_CHAR:
		return loader.LoadClass("[C")
	case AT_SHORT:
		return loader.LoadClass("[S")
	case AT_INT:
		return loader.LoadClass("[I")
	case AT_LONG:
		return loader.LoadClass("[J")
	case AT_FLOAT:
		return loader.LoadClass("[F")
	case AT_DOUBLE:
		return loader.LoadClass("[D")
	default:
		panic("Invalid atype!")
	}
}