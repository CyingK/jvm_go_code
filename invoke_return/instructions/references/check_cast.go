package references

import (
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"jvm_go_code/invoke_return/rtda/heap"
)

type CHECK_CAST struct {
	base.Index16Instruction
}

func (self *CHECK_CAST) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref := stack.PopRef()
	stack.PushRef(ref)
	if ref == nil {
		return
	}
	constantPool := frame.Method().Class().ConstantPool()
	classRef := constantPool.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}

func (self *CHECK_CAST) String() string {
	return "{type：check_cast; " + self.Index16Instruction.String() + "}"
}