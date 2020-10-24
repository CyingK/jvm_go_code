package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

type GET_STATIC struct {
	base.Index16Instruction
}

func (self *GET_STATIC) Execute(frame *rtda.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	fieldRef := constantPool.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.GetClass()
	if !class.GetInitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.GetThread(), class)
		return
	}
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	descriptor := field.GetDescriptor()
	slotId := field.SlotId()
	slots := class.GetStaticVars()
	stack := frame.GetOperandStack()
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
	if self.Index < 10 {
		return "{type：get_static; " + self.Index16Instruction.String() + "}\t"
	}
	return "{type：get_static; " + self.Index16Instruction.String() + "}"
}