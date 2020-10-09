package loads

// 加载指令，从局部变量表获取 int 型变量，然后推入操作数栈顶

/********************
 *    iload			*
 *    iload_0		*
 *    iload_1		*
 *    iload_2		*
 *    iload_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_iload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(value)
}

/*
 * Frame.localVars[ILOAD.index] -> 操作数栈顶
 */
type ILOAD struct {
	base.Index8Instruction
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	util_iload(frame, self.Index)
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	util_iload(frame, 0)
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	util_iload(frame, 1)
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	util_iload(frame, 2)
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	util_iload(frame, 3)
}
