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
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"strconv"
)

func util_dstore(frame *rtda.Frame, index uint) {
	value := frame.GetOperandStack().PopDouble()
	frame.GetLocalVars().SetDouble(index, value)
	if index >= 0 && index <= 3 {
		fmt.Printf("dstore_%d: 从操作数栈弹出 %f 存入局部变量表\n", index, value)
	} else {
		fmt.Printf("dstore: 从操作数栈弹出 %f 存入局部变量表[%d]\n", value, index)
	}
}

/*
 * 操作数栈顶 -> Frame.localVars[index]
 */
type DSTORE struct {
	base.Index8Instruction
}

func (self *DSTORE) Execute(frame *rtda.Frame) {
	util_dstore(frame, self.Index)
}

func (self *DSTORE) String() string {
	return "{type：dstore; Index: " + strconv.Itoa(int(self.Index)) + "}"
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

func (self *DSTORE_0) String() string {
	return "{type：dstore_0; " + self.NoOperandsInstruction.String() + "}"
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

func (self *DSTORE_1) String() string {
	return "{type：dstore_1; " + self.NoOperandsInstruction.String() + "}"
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

func (self *DSTORE_2) String() string {
	return "{type：dstore_2; " + self.NoOperandsInstruction.String() + "}"
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

func (self *DSTORE_3) String() string {
	return "{type：dstore_3; " + self.NoOperandsInstruction.String() + "}"
}