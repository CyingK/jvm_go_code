package loads

// 加载指令，从局部变量表获取 float 型变量，然后推入操作数栈顶

/********************
 *    fload			*
 *    fload_0		*
 *    fload_1		*
 *    fload_2		*
 *    fload_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_fload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(value)
}

/*
 * Frame.localVars[FLOAD.index] -> 操作数栈顶
 */
type FLOAD struct {
	base.Index8Instruction
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	util_fload(frame, self.Index)
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	util_fload(frame, 0)
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	util_fload(frame, 1)
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	util_fload(frame, 2)
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	util_fload(frame, 3)
}
