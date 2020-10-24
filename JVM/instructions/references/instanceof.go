package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

type INSTANCE_OF struct {
	base.Index16Instruction
}

func (self *INSTANCE_OF) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref := stack.PopRef()
	if ref == nil {
		stack.PushInt(0)
		return
	}
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classRef := constantPool.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if ref.IsInstanceOf(class) {
		stack.PushInt(1)
	} else {
		stack.PushInt(0)
	}
}

func (self *INSTANCE_OF) String() string {
	return "{type：instance_of; " + self.Index16Instruction.String() + "}"
}