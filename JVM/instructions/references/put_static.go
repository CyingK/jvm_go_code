package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

type PUT_STATIC struct {
	base.Index16Instruction
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.GetMethod()
	currentClass := currentMethod.GetClass()
	constantPool := currentClass.GetConstantPool()
	fieldRef := constantPool.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.GetClass()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.GetName() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.GetDescriptor()
	slotId := field.SlotId()
	slots := class.GetStaticVars()
	stack := frame.GetOperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		slots.SetInt(slotId, stack.PopInt())
	case 'F':
		slots.SetFloat(slotId, stack.PopFloat())
	case 'J':
		slots.SetLong(slotId, stack.PopLong())
	case 'D':
		slots.SetDouble(slotId, stack.PopDouble())
	case 'L', '[':
		slots.SetRef(slotId, stack.PopRef())
	}
}

func (self *PUT_STATIC) String() string {
	return "{type：put_static; " + self.Index16Instruction.String() + "}"
}