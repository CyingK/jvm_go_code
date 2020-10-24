package stack

// 弹出栈帧

/********************
 *    pop			*
 *    pop2			*
 ********************
 * 2				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	stack.PopSlot()
}

func (self *POP) String() string {
	return "{type：pop; " + self.NoOperandsInstruction.String() + "}\t\t"
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	stack.PopSlot()
	stack.PopSlot()
}

func (self *POP2) String() string {
	return "{type：pop2; " + self.NoOperandsInstruction.String() + "}\t\t"
}