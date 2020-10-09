package stores

// 存储指令，把 long 型变量从操作数栈顶弹出，然后存入局部变量表

/********************
 *    lstore		*
 *    lstore_0		*
 *    lstore_1		*
 *    lstore_2		*
 *    lstore_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_lstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, value)
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type LSTORE struct {
	base.Index8Instruction
}

func (self *LSTORE) Execute(frame *rtda.Frame, index uint) {
	util_lstore(frame, index)
}

/*
 * 操作数栈顶 -> Frame.localVars[0]
 */
type LSTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_0) Execute(frame *rtda.Frame) {
	util_lstore(frame, 0)
}

/*
 * 操作数栈顶 -> Frame.localVars[1]
 */
type LSTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_1) Execute(frame *rtda.Frame) {
	util_lstore(frame, 1)
}

/*
 * 操作数栈顶 -> Frame.localVars[2]
 */
type LSTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_2) Execute(frame *rtda.Frame) {
	util_lstore(frame, 2)
}

/*
 * 操作数栈顶 -> Frame.localVars[3]
 */
type LSTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *LSTORE_3) Execute(frame *rtda.Frame) {
	util_lstore(frame, 3)
}