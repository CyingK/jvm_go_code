package loads

// 加载指令，从局部变量表获取 double 型变量，然后推入操作数栈顶

/********************
 *    dload			*
 *    dload_0		*
 *    dload_1		*
 *    dload_2		*
 *    dload_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_dload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(value)
}

/*
 * Frame.localVars[DLOAD.index] -> 操作数栈顶
 */
type DLOAD struct {
	base.Index8Instruction
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	util_dload(frame, self.Index)
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	util_dload(frame, 0)
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	util_dload(frame, 1)
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	util_dload(frame, 2)
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	util_dload(frame, 3)
}
