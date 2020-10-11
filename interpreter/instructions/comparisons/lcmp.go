package comparisons

// long 比较

/********************
 *    lcmp			*
 ********************
 * 1				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

type LCMP struct {
	base.NoOperandsInstruction
}

func (self *LCMP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v1 > v2 {
		fmt.Printf("lcmp: %v VS %v --> 结果 %d 入操作数栈\n", v1, v2, 1)
		stack.PushInt(1)
	} else if v1 < v2 {
		fmt.Printf("lcmp: %v VS %v --> 结果 %d 入操作数栈\n", v1, v2, -1)
		stack.PushInt(-1)
	} else {
		fmt.Printf("lcmp: %v VS %v --> 结果 %d 入操作数栈\n", v1, v2, 0)
		stack.PushInt(0)
	}
}

func (self *LCMP) String() string {
	return "{type：lcmp; " + self.NoOperandsInstruction.String() + "}"
}
