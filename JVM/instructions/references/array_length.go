package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.GetArrayLength()
	stack.PushInt(arrLen)
}

func (self *ARRAY_LENGTH) String() string {
	return "{type：array_length; " + self.NoOperandsInstruction.String() + "}\t"
}
