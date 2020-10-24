package stack

// 交换栈帧

/********************
 *    swap			*
 ********************
 * 1				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type SWAP struct {
	base.NoOperandsInstruction
}

func (self *SWAP) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot2 := stack.PopSlot()
	slot1 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

func (self *SWAP) String() string {
	return "{type：swap; " + self.NoOperandsInstruction.String() + "}"
}