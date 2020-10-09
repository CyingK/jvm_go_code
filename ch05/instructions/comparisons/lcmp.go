package comparisons

// long 比较

/********************
 *    lcmp			*
 ********************
 * 1				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else {
		stack.PushInt(0)
	}
}
