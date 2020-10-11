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
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

type POP struct {
	base.NoOperandsInstruction
}

func (self *POP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	fmt.Printf("pop: 弹出栈顶元素 %v", slot.String())
}

func (self *POP) String() string {
	return "{type：pop; " + self.NoOperandsInstruction.String() + "}"
}

type POP2 struct {
	base.NoOperandsInstruction
}

func (self *POP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot2 := stack.PopSlot()
	slot1 := stack.PopSlot()
	fmt.Printf("pop2: 弹出栈顶两个元素 %v, %v", slot1.String(), slot2.String())
}

func (self *POP2) String() string {
	return "{type：pop2; " + self.NoOperandsInstruction.String() + "}"
}