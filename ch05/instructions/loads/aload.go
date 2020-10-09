package loads

// 加载指令，从局部变量表获取引用类型变量，然后推入操作数栈顶

/********************
 *    aload			*
 *    aload_0		*
 *    aload_1		*
 *    aload_2		*
 *    aload_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_aload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(value)
}

/*
 * Frame.localVars[ALOAD.index] -> 操作数栈顶
 */
type ALOAD struct {
	base.Index8Instruction
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	util_aload(frame, self.Index)
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	util_aload(frame, 0)
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	util_aload(frame, 1)
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	util_aload(frame, 2)
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	util_aload(frame, 3)
}
