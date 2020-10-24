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
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"strconv"
)

func util_astore(frame *rtda.Frame, index uint) {
	value := frame.GetOperandStack().PopRef()
	frame.GetLocalVars().SetRef(index, value)
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type ASTORE struct {
	base.Index8Instruction
}

func (self *ASTORE) Execute(frame *rtda.Frame) {
	util_astore(frame, self.Index)
}

func (self *ASTORE) String() string {
	return "{type：astore; Index: " + strconv.Itoa(int(self.Index)) + "}\t"
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

func (self *ASTORE_0) String() string {
	return "{type：astore_0; " + self.NoOperandsInstruction.String() + "}\t"
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

func (self *ASTORE_1) String() string {
	return "{type：astore_1; " + self.NoOperandsInstruction.String() + "}\t"
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

func (self *ASTORE_2) String() string {
	return "{type：astore_2; " + self.NoOperandsInstruction.String() + "}\t"
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

func (self *ASTORE_3) String() string {
	return "{type：astore_3; " + self.NoOperandsInstruction.String() + "}\t"
}