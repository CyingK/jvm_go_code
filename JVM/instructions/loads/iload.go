package loads

// 加载指令，从局部变量表获取 int 型变量，然后推入操作数栈顶

/********************
 *    iload			*
 *    iload_0		*
 *    iload_1		*
 *    iload_2		*
 *    iload_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"strconv"
)

func util_iload(frame *rtda.Frame, index uint) {
	value := frame.GetLocalVars().GetInt(index)
	frame.GetOperandStack().PushInt(value)
}

/*
 * Frame.localVars[ILOAD.index] -> 操作数栈顶
 */
type ILOAD struct {
	base.Index8Instruction
}

func (self *ILOAD) Execute(frame *rtda.Frame) {
	util_iload(frame, self.Index)
}

func (self *ILOAD) String() string {
	return "{type：iload; " + strconv.Itoa(int(self.Index)) + "}\t\t"
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type ILOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_0) Execute(frame *rtda.Frame) {
	util_iload(frame, 0)
}

func (self *ILOAD_0) String() string {
	return "{type：iload_0; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type ILOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_1) Execute(frame *rtda.Frame) {
	util_iload(frame, 1)
}

func (self *ILOAD_1) String() string {
	return "{type：iload_1; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type ILOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_2) Execute(frame *rtda.Frame) {
	util_iload(frame, 2)
}

func (self *ILOAD_2) String() string {
	return "{type：iload_2; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type ILOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ILOAD_3) Execute(frame *rtda.Frame) {
	util_iload(frame, 3)
}

func (self *ILOAD_3) String() string {
	return "{type：iload_3; " + self.NoOperandsInstruction.String() + "}\t"
}