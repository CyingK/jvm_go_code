package references

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	fmt.Println("invoke_special: 从操作数栈弹出一个引用")
	frame.OperandStack().PopRef()
}

func (self *INVOKE_SPECIAL) String() string {
	return "{type：invoke_special; " + self.Index16Instruction.String() + "}"
}