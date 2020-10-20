package stack

// 弹出栈帧

/********************
 *    pop			*
 *    pop2			*
 ********************
 * 2				*
 ********************/

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot := stack.PopSlot()
	fmt.Printf("pop: 弹出栈顶元素 %v\n", slot.String())
}

func (self *POP) String() string {
	return "{type：pop; " + self.NoOperandsInstruction.String() + "}"
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	slot2 := stack.PopSlot()
	slot1 := stack.PopSlot()
	fmt.Printf("pop2: 弹出栈顶两个元素 %v, %v\n", slot1.String(), slot2.String())
}

func (self *POP2) String() string {
	return "{type：pop2; " + self.NoOperandsInstruction.String() + "}"
}