package stores

// 存储指令，把 int 型变量从操作数栈顶弹出，然后存入局部变量表

/********************
 *    istore		*
 *    istore_0		*
 *    istore_1		*
 *    istore_2		*
 *    istore_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_istore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, value)
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type ISTORE struct {
	base.Index8Instruction
}

func (self *ISTORE) Execute(frame *rtda.Frame, index uint) {
	util_istore(frame, index)
}

/*
 * 操作数栈顶 -> Frame.localVars[0]
 */
type ISTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_0) Execute(frame *rtda.Frame) {
	util_istore(frame, 0)
}

/*
 * 操作数栈顶 -> Frame.localVars[1]
 */
type ISTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_1) Execute(frame *rtda.Frame) {
	util_istore(frame, 1)
}

/*
 * 操作数栈顶 -> Frame.localVars[2]
 */
type ISTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_2) Execute(frame *rtda.Frame) {
	util_istore(frame, 2)
}

/*
 * 操作数栈顶 -> Frame.localVars[3]
 */
type ISTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ISTORE_3) Execute(frame *rtda.Frame) {
	util_istore(frame, 3)
}