package comparisons

// long 比较

/********************
 *    lcmp			*
 ********************
 * 1				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
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

func (self *LCMP) String() string {
	return "{type：lcmp; " + self.NoOperandsInstruction.String() + "}\t\t"
}
