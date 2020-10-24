package math

// 按位异或运算

/********************
 *    ixor			*
 *    lxor			*
 ********************
 * 2				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * int 按位异或
 */
type IXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	stack.PushInt(result)
}

func (self *IXOR) String() string {
	return "{type：ixor; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * long 按位异或
 */
type LXOR struct {
	base.NoOperandsInstruction
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 ^ v2
	stack.PushLong(result)
}

func (self *LXOR) String() string {
	return "{type：lxor; " + self.NoOperandsInstruction.String() + "}\t\t"
}