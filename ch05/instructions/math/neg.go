package math

// 取反运算

/********************
 *    dneg			*
 *    fneg			*
 *    ineg			*
 *    lneg			*
 ********************
 * 4				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

/*
 * double 取反
 */
type DNEG struct {
	base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopDouble()
	stack.PushDouble(-value)
}

/*
 * float 取反
 */
type FNEG struct {
	base.NoOperandsInstruction
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopFloat()
	stack.PushFloat(-value)
}

/*
 * long 取反
 */
type LNEG struct {
	base.NoOperandsInstruction
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopLong()
	stack.PushLong(-value)
}

/*
 * int 取反
 */
type INEG struct {
	base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	stack.PushInt(-value)
}