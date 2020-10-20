package loads

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
)

func checkNoNil(ref *heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func checkIndex(arrLen int, index int32) {
	if index < 0 || index >= int32(arrLen) {
		panic("ArrayIndexOutOfBoundsException")
	}
}

type AALOAD struct {
	base.NoOperandsInstruction
}

func (self *AALOAD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNoNil(arrRef)
	refs := arrRef.GetRefs()
	checkIndex(len(refs), index)
	stack.PushRef(refs[index])
	fmt.Printf("aaload: 将下标 %v 的元素 %v 推入操作数栈\n", index, refs[index])
}

func (self *AALOAD) String() string {
	return "{type：aaload; " + self.NoOperandsInstruction.String() + "}"
}

type BALOAD struct {
	base.NoOperandsInstruction
}

func (self *BALOAD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNoNil(arrRef)
	bytes := arrRef.GetBytes()
	checkIndex(len(bytes), index)
	stack.PushInt(int32(bytes[index]))
	fmt.Printf("baload: 将下标 %v 的元素 %v 推入操作数栈\n", index, bytes[index])
}

func (self *BALOAD) String() string {
	return "{type：baload; " + self.NoOperandsInstruction.String() + "}"
}

type CALOAD struct {
	base.NoOperandsInstruction
}

func (self *CALOAD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNoNil(arrRef)
	chars := arrRef.GetChars()
	checkIndex(len(chars), index)
	stack.PushInt(int32(chars[index]))
	fmt.Printf("caload: 将下标 %v 的元素 %v 推入操作数栈\n", index, chars[index])
}

func (self *CALOAD) String() string {
	return "{type：caload; " + self.NoOperandsInstruction.String() + "}"
}

type DALOAD struct {
	base.NoOperandsInstruction
}

func (self *DALOAD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNoNil(arrRef)
	doubles := arrRef.GetDoubles()
	checkIndex(len(doubles), index)
	stack.PushDouble(doubles[index])
	fmt.Printf("daload: 将下标 %v 的元素 %v 推入操作数栈\n", index, doubles[index])
}

func (self *DALOAD) String() string {
	return "{type：daload; " + self.NoOperandsInstruction.String() + "}"
}

type FALOAD struct {
	base.NoOperandsInstruction
}

func (self *FALOAD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNoNil(arrRef)
	floats := arrRef.GetFloats()
	checkIndex(len(floats), index)
	stack.PushFloat(floats[index])
	fmt.Printf("faload: 将下标 %v 的元素 %v 推入操作数栈\n", index, floats[index])
}

func (self *FALOAD) String() string {
	return "{type：faload; " + self.NoOperandsInstruction.String() + "}"
}

type IALOAD struct {
	base.NoOperandsInstruction
}

func (self *IALOAD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNoNil(arrRef)
	ints := arrRef.GetInts()
	checkIndex(len(ints), index)
	stack.PushInt(ints[index])
	fmt.Printf("iaload: 将下标 %v 的元素 %v 推入操作数栈\n", index, ints[index])
}

func (self *IALOAD) String() string {
	return "{type：iaload; " + self.NoOperandsInstruction.String() + "}"
}

type LALOAD struct {
	base.NoOperandsInstruction
}

func (self *LALOAD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNoNil(arrRef)
	longs := arrRef.GetLongs()
	checkIndex(len(longs), index)
	stack.PushLong(longs[index])
	fmt.Printf("laload: 将下标 %v 的元素 %v 推入操作数栈\n", index, longs[index])
}

func (self *LALOAD) String() string {
	return "{type：laload; " + self.NoOperandsInstruction.String() + "}"
}

type SALOAD struct {
	base.NoOperandsInstruction
}

func (self *SALOAD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	index := stack.PopInt()
	arrRef := stack.PopRef()
	checkNoNil(arrRef)
	shorts := arrRef.GetShorts()
	checkIndex(len(shorts), index)
	stack.PushInt(int32(shorts[index]))
	fmt.Printf("saload: 将下标 %v 的元素 %v 推入操作数栈\n", index, shorts[index])
}

func (self *SALOAD) String() string {
	return "{type：saload; " + self.NoOperandsInstruction.String() + "}"
}