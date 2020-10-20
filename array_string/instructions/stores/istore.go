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
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"strconv"
)

func util_istore(frame *rtda.Frame, index uint) {
	value := frame.GetOperandStack().PopInt()
	frame.GetLocalVars().SetInt(index, value)
	if index >= 0 && index <= 3 {
		fmt.Printf("istore_%d: 从操作数栈弹出 %d 存入局部变量表\n", index, value)
	} else {
		fmt.Printf("istore: 从操作数栈弹出 %d 存入局部变量表[%d]\n", value, index)
	}
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type ISTORE struct {
	base.Index8Instruction
}

func (self *ISTORE) Execute(frame *rtda.Frame) {
	util_istore(frame, self.Index)
}

func (self *ISTORE) String() string {
	return "{type：istore; Index: " + strconv.Itoa(int(self.Index)) + "}"
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

func (self *ISTORE_0) String() string {
	return "{type：istore_0; " + self.NoOperandsInstruction.String() + "}"
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

func (self *ISTORE_1) String() string {
	return "{type：istore_1; " + self.NoOperandsInstruction.String() + "}"
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

func (self *ISTORE_2) String() string {
	return "{type：istore_2; " + self.NoOperandsInstruction.String() + "}"
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

func (self *ISTORE_3) String() string {
	return "{type：istore_3; " + self.NoOperandsInstruction.String() + "}"
}