package stores

// 存储指令，把 double 型变量从操作数栈顶弹出，然后存入局部变量表

/********************
 *    dstore		*
 *    dstore_0		*
 *    dstore_1		*
 *    dstore_2		*
 *    dstore_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_dstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, value)
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type DSTORE struct {
	base.Index8Instruction
}

func (self *DSTORE) Execute(frame *rtda.Frame, index uint) {
	util_dstore(frame, index)
}

/*
 * 操作数栈顶 -> Frame.localVars[0]
 */
type DSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_0) Execute(frame *rtda.Frame) {
	util_dstore(frame, 0)
}

/*
 * 操作数栈顶 -> Frame.localVars[1]
 */
type DSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_1) Execute(frame *rtda.Frame) {
	util_dstore(frame, 1)
}

/*
 * 操作数栈顶 -> Frame.localVars[2]
 */
type DSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_2) Execute(frame *rtda.Frame) {
	util_dstore(frame, 2)
}

/*
 * 操作数栈顶 -> Frame.localVars[3]
 */
type DSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *DSTORE_3) Execute(frame *rtda.Frame) {
	util_dstore(frame, 3)
}