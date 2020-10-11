package loads

// 加载指令，从局部变量表获取引用类型变量，然后推入操作数栈顶

/********************
 *    aload			*
 *    aload_0		*
 *    aload_1		*
 *    aload_2		*
 *    aload_3		*
 ********************
 * 5				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
	"strconv"
)

func util_aload(frame *rtda.Frame, index uint) {
	value := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(value)
	if index >= 0 && index <= 3 {
		fmt.Printf("aload_%d: 从局部变量表[%d]取出 %s 压入操作数栈\n", index, index, value.String())
	} else {
		fmt.Printf("aload: 从局部变量表取出 %s 压入操作数栈\n", value.String())
	}
}

/*
 * Frame.localVars[ALOAD.index] -> 操作数栈顶
 */
type ALOAD struct {
	base.Index8Instruction
}

func (self *ALOAD) Execute(frame *rtda.Frame) {
	util_aload(frame, self.Index)
}

func (self *ALOAD) String() string {
	return "{type：aload; " + strconv.Itoa(int(self.Index)) + "}"
}

/*
 * Frame.localVars[0] -> 操作数栈顶
 */
type ALOAD_0 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_0) Execute(frame *rtda.Frame) {
	util_aload(frame, 0)
}

func (self *ALOAD_0) String() string {
	return "{type：aload_0; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * Frame.localVars[1] -> 操作数栈顶
 */
type ALOAD_1 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_1) Execute(frame *rtda.Frame) {
	util_aload(frame, 1)
}

func (self *ALOAD_1) String() string {
	return "{type：aload_1; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * Frame.localVars[2] -> 操作数栈顶
 */
type ALOAD_2 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_2) Execute(frame *rtda.Frame) {
	util_aload(frame, 2)
}

func (self *ALOAD_2) String() string {
	return "{type：aload_2; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * Frame.localVars[3] -> 操作数栈顶
 */
type ALOAD_3 struct {
	base.NoOperandsInstruction
}

func (self *ALOAD_3) Execute(frame *rtda.Frame) {
	util_aload(frame, 3)
}

func (self *ALOAD_3) String() string {
	return "{type：aload_3; " + self.NoOperandsInstruction.String() + "}"
}
