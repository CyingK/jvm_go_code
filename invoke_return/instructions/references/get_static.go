package references

import (
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"jvm_go_code/invoke_return/rtda/heap"
)

type GET_STATIC struct {
	base.Index16Instruction
}

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	fieldRef := constantPool.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
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
	}
}

func (self *GET_STATIC) String() string {
	return "{type：get_static; " + self.Index16Instruction.String() + "}"
}