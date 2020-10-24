package stores

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
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
}

func (self *AASTORE) String() string {
	return "{type：aastore; " + self.NoOperandsInstruction.String() + "}\t"
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
}

func (self *BASTORE) String() string {
	return "{type：bastore; " + self.NoOperandsInstruction.String() + "}\t"
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
}

func (self *CASTORE) String() string {
	return "{type：castore; " + self.NoOperandsInstruction.String() + "}\t"
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
}

func (self *DASTORE) String() string {
	return "{type：dastore; " + self.NoOperandsInstruction.String() + "}\t"
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
}

func (self *FASTORE) String() string {
	return "{type：fastore; " + self.NoOperandsInstruction.String() + "}\t"
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
}

func (self *IASTORE) String() string {
	return "{type：iastore; " + self.NoOperandsInstruction.String() + "}\t"
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
}

func (self *LASTORE) String() string {
	return "{type：lastore; " + self.NoOperandsInstruction.String() + "}\t"
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
}

func (self *SASTORE) String() string {
	return "{type：sastore; " + self.NoOperandsInstruction.String() + "}\t"
}