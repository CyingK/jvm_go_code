package references

import (
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"jvm_go_code/invoke_return/rtda/heap"
	"log"
)

type PUT_STATIC struct {
	base.Index16Instruction
}

func (self *PUT_STATIC) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	constantPool := currentClass.ConstantPool()
	fieldRef := constantPool.GetConstant(self.Index).(*heap.FieldRef)
	field := fieldRef.ResolvedField()
	class := field.Class()
	if !field.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	if field.IsFinal() {
		if currentClass != class || currentMethod.Name() != "<clinit>" {
			panic("java.lang.IllegalAccessError")
		}
	}
	descriptor := field.Descriptor()
	slotId := field.SlotId()
	slots := class.StaticVars()
	stack := frame.OperandStack()
	switch descriptor[0] {
	case 'Z', 'B', 'C', 'S', 'I':
		log.Println(len(class.StaticVars()), class.GetName())
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