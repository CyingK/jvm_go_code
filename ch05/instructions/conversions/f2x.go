package conversions

// 把float变量强制转换成其他类型

/********************
 *    f2d			*
 *    f2i			*
 *    f2l			*
 ********************
 * 3				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

/*
 * float -> double
 */
type F2D struct {
	base.NoOperandsInstruction
}

func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_float := stack.PopFloat()
	_double := float64(_float)
	stack.PushDouble(_double)
}

/*
 * float -> int
 */
type F2I struct {
	base.NoOperandsInstruction
}

func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_float := stack.PopFloat()
	_int := int32(_float)
	stack.PushInt(_int)
}

/*
 * float -> long
 */
type F2L struct {
	base.NoOperandsInstruction
}

func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_float := stack.PopFloat()
	_long := int64(_float)
	stack.PushLong(_long)
}

