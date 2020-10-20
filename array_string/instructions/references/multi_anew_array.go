package references

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
	"strconv"
)

type MULTI_ANEW_ARRAY struct {
	index		uint16
	dimensions	uint8
}

func (self *MULTI_ANEW_ARRAY) GetOperands(reader *base.ByteCodeReader) {
	self.index = reader.ReadUint16()
	self.dimensions = reader.ReadUint8()
}

func (self *MULTI_ANEW_ARRAY) Execute(frame *rtda.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classRef := constantPool.GetConstant(uint(self.index)).(*heap.ClassRef)
	arrClass := classRef.ResolvedClass()
	stack := frame.GetOperandStack()
	counts := popAndCheckCounts(stack, int(self.dimensions))
	arr := newMultiDimensionalArray(counts, arrClass)
	stack.PushRef(arr)
	fmt.Println("multi_anew_array: 创建多维数组")
}

func (self *MULTI_ANEW_ARRAY) String() string {
	return "{type：multi_anew_array; index: " + strconv.Itoa(int(self.index)) + ", dimensions: " + strconv.Itoa(int(self.dimensions)) + "}"
}

func newMultiDimensionalArray(counts []int32, arrClass *heap.Class) *heap.Object {
	count := uint(counts[0])
	arr := arrClass.NewArray(count)
	if len(counts) > 1 {
		refs := arr.GetRefs()
		for item := range refs {
			refs[item] = newMultiDimensionalArray(counts[1:], arrClass.GetComponentClass())
		}
	}
	return arr
}

func popAndCheckCounts(stack *rtda.OperandStack, dimensions int) []int32 {
	counts := make([]int32, dimensions)
	for index := dimensions - 1; index >= 0; index-- {
		counts[index] = stack.PopInt()
		if counts[index] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}