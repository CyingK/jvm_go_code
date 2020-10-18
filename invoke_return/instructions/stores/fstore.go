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
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"strconv"
)

func util_fstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, value)
	if index >= 0 && index <= 3 {
		fmt.Printf("fstore_%d: 从操作数栈弹出 %f 存入局部变量表\n", index, value)
	} else {
		fmt.Printf("fstore: 从操作数栈弹出 %f 存入局部变量表[%d]\n", value, index)
	}
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type FSTORE struct {
	base.Index8Instruction
}

func (self *FSTORE) Execute(frame *rtda.Frame) {
	util_fstore(frame, self.Index)
}

func (self *FSTORE) String() string {
	return "{type：fstore; Index: " + strconv.Itoa(int(self.Index)) + "}"
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

func (self *FSTORE_0) String() string {
	return "{type：fstore_0; " + self.NoOperandsInstruction.String() + "}"
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

func (self *FSTORE_1) String() string {
	return "{type：fstore_1; " + self.NoOperandsInstruction.String() + "}"
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

func (self *FSTORE_2) String() string {
	return "{type：fstore_2; " + self.NoOperandsInstruction.String() + "}"
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

func (self *FSTORE_3) String() string {
	return "{type：fstore_3; " + self.NoOperandsInstruction.String() + "}"
}