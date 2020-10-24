package loads

// 加载指令，从局部变量表获取 float 型变量，然后推入操作数栈顶

/********************
 *    fload			*
 *    fload_0		*
 *    fload_1		*
 *    fload_2		*
 *    fload_3		*
 ********************
 * 5				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"strconv"
)

func util_fload(frame *rtda.Frame, index uint) {
	value := frame.GetLocalVars().GetFloat(index)
	frame.GetOperandStack().PushFloat(value)
}

/*
 * Frame.localVars[FLOAD.index] -> 操作数栈顶
 */
type FLOAD struct {
	base.Index8Instruction
}

func (self *FLOAD) Execute(frame *rtda.Frame) {
	util_fload(frame, self.Index)
}

func (self *FLOAD) String() string {
	return "{type：fload; " + strconv.Itoa(int(self.Index)) + "}\t"
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type FLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_0) Execute(frame *rtda.Frame) {
	util_fload(frame, 0)
}

func (self *FLOAD_0) String() string {
	return "{type：fload_0; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type FLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_1) Execute(frame *rtda.Frame) {
	util_fload(frame, 1)
}

func (self *FLOAD_1) String() string {
	return "{type：fload_1; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type FLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_2) Execute(frame *rtda.Frame) {
	util_fload(frame, 2)
}

func (self *FLOAD_2) String() string {
	return "{type：fload_2; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type FLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *FLOAD_3) Execute(frame *rtda.Frame) {
	util_fload(frame, 3)
}

func (self *FLOAD_3) String() string {
	return "{type：fload_3; " + self.NoOperandsInstruction.String() + "}\t"
}