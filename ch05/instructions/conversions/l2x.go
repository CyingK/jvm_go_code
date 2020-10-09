package conversions

// 把 long 变量强制转换成其他类型

/********************
 *    d2f			*
 *    d2i			*
 *    d2l			*
 ********************
 * 3				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

/*
 * long -> float
 */
type L2F struct {
	base.NoOperandsInstruction
}

func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_long := stack.PopLong()
	_float := float32(_long)
	stack.PushFloat(_float)
}

/*
 * long -> int
 */
type L2I struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_long := stack.PopLong()
	_int := int32(_long)
	stack.PushInt(_int)
}

/*
 * long -> double
 */
type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_long := stack.PopLong()
	_double := float64(_long)
	stack.PushDouble(_double)
}

