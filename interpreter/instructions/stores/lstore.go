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
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
	"strconv"
)

func util_lstore(frame *rtda.Frame, index uint) {
	value := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, value)
	if index >= 0 && index <= 3 {
		fmt.Printf("lstore_%d: 从操作数栈弹出 %d 存入局部变量表\n", index, value)
	} else {
		fmt.Printf("lstore: 从操作数栈弹出 %d 存入局部变量表[%d]\n", value, index)
	}
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type LSTORE struct {
	base.Index8Instruction
}

func (self *LSTORE) Execute(frame *rtda.Frame) {
	util_lstore(frame, self.Index)
}

func (self *LSTORE) String() string {
	return "{type：lstore; Index: " + strconv.Itoa(int(self.Index)) + "}"
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

func (self *LSTORE_0) String() string {
	return "{type：lstore_0; " + self.NoOperandsInstruction.String() + "}"
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

func (self *LSTORE_1) String() string {
	return "{type：lstore_1; " + self.NoOperandsInstruction.String() + "}"
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

func (self *LSTORE_2) String() string {
	return "{type：lstore_2; " + self.NoOperandsInstruction.String() + "}"
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

func (self *LSTORE_3) String() string {
	return "{type：lstore_3; " + self.NoOperandsInstruction.String() + "}"
}