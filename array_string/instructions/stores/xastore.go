package stores

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
)

func checkNotNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

type AASTORE struct {
	base.NoOperandsInstruction
}

func (self *AASTORE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref := stack.PopRef()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	refs := arrRef.GetRefs()
	checkIndex(len(refs), index)
	refs[index] = ref
	fmt.Printf("aastore: 从操作数栈取出 %v, 放入 %v 位置\n", ref.String(), index)
}

func (self *AASTORE) String() string {
	return "{type：aastore; " + self.NoOperandsInstruction.String() + "}"
}

type BASTORE struct {
	base.NoOperandsInstruction
}

func (self *BASTORE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	bytes := arrRef.GetBytes()
	checkIndex(len(bytes), index)
	bytes[index] = int8(value)
	fmt.Printf("bastore: 从操作数栈取出 %v, 放入 %v 位置\n", value, index)
}

func (self *BASTORE) String() string {
	return "{type：bastore; " + self.NoOperandsInstruction.String() + "}"
}

type CASTORE struct {
	base.NoOperandsInstruction
}

func (self *CASTORE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	chars := arrRef.GetChars()
	checkIndex(len(chars), index)
	chars[index] = uint16(value)
	fmt.Printf("castore: 从操作数栈取出 %v, 放入 %v 位置\n", value, index)
}

func (self *CASTORE) String() string {
	return "{type：castore; " + self.NoOperandsInstruction.String() + "}"
}

type DASTORE struct {
	base.NoOperandsInstruction
}

func (self *DASTORE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopDouble()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	doubles := arrRef.GetDoubles()
	checkIndex(len(doubles), index)
	doubles[index] = value
	fmt.Printf("dastore: 从操作数栈取出 %v, 放入 %v 位置\n", value, index)
}

func (self *DASTORE) String() string {
	return "{type：dastore; " + self.NoOperandsInstruction.String() + "}"
}

type FASTORE struct {
	base.NoOperandsInstruction
}

func (self *FASTORE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopFloat()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	floats := arrRef.GetFloats()
	checkIndex(len(floats), index)
	floats[index] = value
	fmt.Printf("fastore: 从操作数栈取出 %v, 放入 %v 位置\n", value, index)
}

func (self *FASTORE) String() string {
	return "{type：fastore; " + self.NoOperandsInstruction.String() + "}"
}

type IASTORE struct {
	base.NoOperandsInstruction
}

func (self *IASTORE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	ints := arrRef.GetInts()
	checkIndex(len(ints), index)
	ints[index] = value
	fmt.Printf("iastore: 从操作数栈取出 %v, 放入 %v 位置\n", value, index)
}

func (self *IASTORE) String() string {
	return "{type：iastore; " + self.NoOperandsInstruction.String() + "}"
}

type LASTORE struct {
	base.NoOperandsInstruction
}

func (self *LASTORE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopLong()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	longs := arrRef.GetLongs()
	checkIndex(len(longs), index)
	longs[index] = value
	fmt.Printf("lastore: 从操作数栈取出 %v, 放入 %v 位置\n", value, index)
}

func (self *LASTORE) String() string {
	return "{type：lastore; " + self.NoOperandsInstruction.String() + "}"
}

type SASTORE struct {
	base.NoOperandsInstruction
}

func (self *SASTORE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopInt()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNotNil(arrRef)
	shorts := arrRef.GetShorts()
	checkIndex(len(shorts), index)
	shorts[index] = int16(value)
	fmt.Printf("sastore: 从操作数栈取出 %v, 放入 %v 位置\n", value, index)
}

func (self *SASTORE) String() string {
	return "{type：sastore; " + self.NoOperandsInstruction.String() + "}"
}