package conversions

// 把int变量强制转换成其他类型

/********************
 *    d2b			*
 *    d2c			*
 *    d2s			*
 *    d2f			*
 *    d2i			*
 *    d2l			*
 ********************
 * 6				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

/*
 * int -> byte
 */
type I2B struct {
	base.NoOperandsInstruction
}

func (self *I2B) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_int := stack.PopInt()
	_byte := int32(int8(_int))
	stack.PushInt(_byte)
}

/*
 * int -> char
 */
type I2C struct {
	base.NoOperandsInstruction
}

func (self *I2C) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_int := stack.PopInt()
	_char := int32(uint16(_int))
	stack.PushInt(_char)
}

/*
 * int -> short
 */
type I2S struct {
	base.NoOperandsInstruction
}

func (self *I2S) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_int := stack.PopInt()
	_short := int32(int16(_int))
	stack.PushInt(_short)
}

/*
 * int -> float
 */
type I2F struct {
	base.NoOperandsInstruction
}

func (self *I2F) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_int := stack.PopInt()
	_float := float32(_int)
	stack.PushFloat(_float)
}

/*
 * int -> double
 */
type I2D struct {
	base.NoOperandsInstruction
}

func (self *I2D) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_int := stack.PopInt()
	_double := float64(_int)
	stack.PushDouble(_double)
}

/*
 * int -> long
 */
type I2L struct {
	base.NoOperandsInstruction
}

func (self *I2L) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	_int := stack.PopInt()
	_long := int64(_int)
	stack.PushLong(_long)
}

