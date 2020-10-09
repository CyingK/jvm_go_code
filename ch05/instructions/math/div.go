package math

// 除法运算

/********************
 *    ddiv			*
 *    fdiv			*
 *    idiv			*
 *    ldiv			*
 ********************
 * 4				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

/*
 * double 相除
 */
type DDIV struct {
	base.NoOperandsInstruction
}

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	stack.PushDouble(result)
}

/*
 * float 相除
 */
type FDIV struct {
	base.NoOperandsInstruction
}

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	stack.PushFloat(result)
}

/*
 * long 相除
 */
type LDIV struct {
	base.NoOperandsInstruction
}

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushLong(result)
}

/*
 * int 相除
 */
type IDIV struct {
	base.NoOperandsInstruction
}

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	stack.PushInt(result)
}