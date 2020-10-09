package conversions

// 把double变量强制转换成其他类型

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
 * double -> float
 */
type D2F struct {
	base.NoOperandsInstruction
}

func (self *D2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_double := stack.PopDouble()
	_float := float32(_double)
	stack.PushFloat(_float)
}

/*
 * double -> int
 */
type D2I struct {
	base.NoOperandsInstruction
}

func (self *D2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_double := stack.PopDouble()
	_int := int32(_double)
	stack.PushInt(_int)
}

/*
 * double -> long
 */
type D2L struct {
	base.NoOperandsInstruction
}

func (self *D2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_double := stack.PopDouble()
	_long := int64(_double)
	stack.PushLong(_long)
}

