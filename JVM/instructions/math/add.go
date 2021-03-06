package math

// 加法运算

/********************
 *    dadd			*
 *    fadd			*
 *    iadd			*
 *    ladd			*
 ********************
 * 4				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * double 相加
 */
type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 + v2
	stack.PushDouble(result)
}

func (self *DADD) String() string {
	return "{type：dadd; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * float 相加
 */
type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	stack.PushFloat(result)
}

func (self *FADD) String() string {
	return "{type：fadd; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * long 相加
 */
type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	stack.PushLong(result)
}

func (self *LADD) String() string {
	return "{type：ladd; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * int 相加
 */
type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	stack.PushInt(result)
}

func (self *IADD) String() string {
	return "{type：ladd; " + self.NoOperandsInstruction.String() + "}\t\t"
}