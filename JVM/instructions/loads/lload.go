package loads

// 加载指令，从局部变量表获取变量，然后推入操作数栈顶

/********************
 *    lload			*
 *    lload_0		*
 *    lload_1		*
 *    lload_2		*
 *    lload_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"strconv"
)

func util_lload(frame *rtda.Frame, index uint) {
	value := frame.GetLocalVars().GetLong(index)
	frame.GetOperandStack().PushLong(value)
}

/*
 * Frame.localVars[LLOAD.index] -> 操作数栈顶
 */
type LLOAD struct {
	base.Index8Instruction
}

func (self *LLOAD) Execute(frame *rtda.Frame) {
	util_lload(frame, self.Index)
}

func (self *LLOAD) String() string {
	return "{type：lload; " + strconv.Itoa(int(self.Index)) + "}\t\t"
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type LLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_0) Execute(frame *rtda.Frame) {
	util_lload(frame, 0)
}

func (self *LLOAD_0) String() string {
	return "{type：lload_0; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type LLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_1) Execute(frame *rtda.Frame) {
	util_lload(frame, 1)
}

func (self *LLOAD_1) String() string {
	return "{type：lload_1; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type LLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_2) Execute(frame *rtda.Frame) {
	util_lload(frame, 2)
}

func (self *LLOAD_2) String() string {
	return "{type：lload_2; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type LLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *LLOAD_3) Execute(frame *rtda.Frame) {
	util_lload(frame, 3)
}

func (self *LLOAD_3) String() string {
	return "{type：lload_3; " + self.NoOperandsInstruction.String() + "}\t"
}