package stores

// 存储指令，把引用变量从操作数栈顶弹出，然后存入局部变量表

/********************
 *    astore		*
 *    astore_0		*
 *    astore_1		*
 *    astore_2		*
 *    astore_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_astore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, value)
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type ASTORE struct {
	base.Index8Instruction
}

func (self *ASTORE) Execute(frame *rtda.Frame, index uint) {
	util_astore(frame, index)
}

/*
 * 操作数栈顶 -> Frame.localVars[0]
 */
type ASTORE_0 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_0) Execute(frame *rtda.Frame) {
	util_astore(frame, 0)
}

/*
 * 操作数栈顶 -> Frame.localVars[1]
 */
type ASTORE_1 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_1) Execute(frame *rtda.Frame) {
	util_astore(frame, 1)
}

/*
 * 操作数栈顶 -> Frame.localVars[2]
 */
type ASTORE_2 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_2) Execute(frame *rtda.Frame) {
	util_astore(frame, 2)
}

/*
 * 操作数栈顶 -> Frame.localVars[3]
 */
type ASTORE_3 struct {
	base.NoOperandsInstruction
}

func (self *ASTORE_3) Execute(frame *rtda.Frame) {
	util_astore(frame, 3)
}