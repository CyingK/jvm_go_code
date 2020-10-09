package loads

// 加载指令，从局部变量表获取变量，然后推入操作数栈顶

/********************
 *    lload			*
 *    lload_0		*
 *    lload_1		*
 *    lload_2		*
 *    lload_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_lload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(value)
}

/*
 * Frame.localVars[LLOAD.index] -> 操作数栈顶
 */
type LLOAD struct {
	base.Index8Instruction
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	util_lload(frame, self.Index)
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	util_lload(frame, 0)
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	util_lload(frame, 1)
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	util_lload(frame, 2)
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	util_lload(frame, 3)
}
