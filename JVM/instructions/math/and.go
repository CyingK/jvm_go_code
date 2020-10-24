package math

// 按位与运算

/********************
 *    iand			*
 *    land			*
 ********************
 * 2				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * int 按位与
 */
type IAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	stack.PushInt(result)
}

func (self *IAND) String() string {
	return "{type：iand; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * long 按位与
 */
type LAND struct {
	base.NoOperandsInstruction
}

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	stack.PushLong(result)
}

func (self *LAND) String() string {
	return "{type：land; " + self.NoOperandsInstruction.String() + "}\t\t"
}