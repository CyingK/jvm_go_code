package stores

// 存储指令，把 float 型变量从操作数栈顶弹出，然后存入局部变量表

/********************
 *    fstore		*
 *    fstore_0		*
 *    fstore_1		*
 *    fstore_2		*
 *    fstore_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_fstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, value)
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type FSTORE struct {
	base.Index8Instruction
}

func (self *FSTORE) Execute(frame *rtda.Frame, index uint) {
	util_fstore(frame, index)
}

/*
 * 操作数栈顶 -> Frame.localVars[0]
 */
type FSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_0) Execute(frame *rtda.Frame) {
	util_fstore(frame, 0)
}

/*
 * 操作数栈顶 -> Frame.localVars[1]
 */
type FSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_1) Execute(frame *rtda.Frame) {
	util_fstore(frame, 1)
}

/*
 * 操作数栈顶 -> Frame.localVars[2]
 */
type FSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_2) Execute(frame *rtda.Frame) {
	util_fstore(frame, 2)
}

/*
 * 操作数栈顶 -> Frame.localVars[3]
 */
type FSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *FSTORE_3) Execute(frame *rtda.Frame) {
	util_fstore(frame, 3)
}