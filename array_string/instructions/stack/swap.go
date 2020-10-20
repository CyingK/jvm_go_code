package stack

// 交换栈帧

/********************
 *    swap			*
 ********************
 * 1				*
 ********************/

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
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
	fmt.Printf("swap: 交换栈顶两个元素 [%v, %v] -> [%v, %v]\n",
		slot1.String(), slot2.String(),
		slot2.String(), slot1.String())
}

func (self *SWAP) String() string {
	return "{type：swap; " + self.NoOperandsInstruction.String() + "}"
}