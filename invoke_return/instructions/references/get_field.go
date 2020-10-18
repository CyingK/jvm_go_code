package references

import (
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"jvm_go_code/invoke_return/rtda/heap"
)

type GET_FIELD struct {
	base.Index16Instruction
}

func (self *GET_FIELD) Execute(frame *rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	fieldRef := constantPool.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	if field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	stack := frame.OperandStack()
	ref := stack.PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := ref.Fields()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		stack.PushInt(slots.GetInt(slotId))
	case 'F':
		stack.PushFloat(slots.GetFloat(slotId))
	case 'J':
		stack.PushLong(slots.GetLong(slotId))
	case 'D':
		stack.PushDouble(slots.GetDouble(slotId))
	case 'L', '[':
		stack.PushRef(slots.GetRef(slotId))
	default:
	}
}

func (self *GET_FIELD) String() string {
	return "{type：get_field; " + self.Index16Instruction.String() + "}"
}