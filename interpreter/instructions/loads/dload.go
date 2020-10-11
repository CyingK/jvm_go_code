package loads

// 加载指令，从局部变量表获取 double 型变量，然后推入操作数栈顶

/********************
 *    dload			*
 *    dload_0		*
 *    dload_1		*
 *    dload_2		*
 *    dload_3		*
 ********************
 * 5				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
	"strconv"
)

func util_dload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(value)
	if index >= 0 && index <= 3 {
		fmt.Printf("dload_%d: 从局部变量表[%d]取出 %f 压入操作数栈\n", index, index, value)
	} else {
		fmt.Printf("dload: 从局部变量表取出 %f 压入操作数栈\n", value)
	}
}

/*
 * Frame.localVars[DLOAD.index] -> 操作数栈顶
 */
type DLOAD struct {
	base.Index8Instruction
}

func (self *DLOAD) Execute(frame *rtda.Frame) {
	util_dload(frame, self.Index)
}

func (self *DLOAD) String() string {
	return "{type：dload; " + strconv.Itoa(int(self.Index)) + "}"
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type DLOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_0) Execute(frame *rtda.Frame) {
	util_dload(frame, 0)
}

func (self *DLOAD_0) String() string {
	return "{type：dload_0; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type DLOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_1) Execute(frame *rtda.Frame) {
	util_dload(frame, 1)
}

func (self *DLOAD_1) String() string {
	return "{type：dload_1; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type DLOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_2) Execute(frame *rtda.Frame) {
	util_dload(frame, 2)
}

func (self *DLOAD_2) String() string {
	return "{type：dload_2; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type DLOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *DLOAD_3) Execute(frame *rtda.Frame) {
	util_dload(frame, 3)
}

func (self *DLOAD_3) String() string {
	return "{type：dload_3; " + self.NoOperandsInstruction.String() + "}"
}
