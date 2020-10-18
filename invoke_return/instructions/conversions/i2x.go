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
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
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
	fmt.Printf("%v<int> 转 %v<byte>, 然后放入操作数栈\n", _int, _byte)
	stack.PushInt(_byte)
}

func (self *I2B) String() string {
	return "{type：i2b; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<int> 转 %v<char>, 然后放入操作数栈\n", _int, _char)
	stack.PushInt(_char)
}

func (self *I2C) String() string {
	return "{type：i2c; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<int> 转 %v<short>, 然后放入操作数栈\n", _int, _short)
	stack.PushInt(_short)
}

func (self *I2S) String() string {
	return "{type：i2s; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<int> 转 %v<float>, 然后放入操作数栈\n", _int, _float)
	stack.PushFloat(_float)
}

func (self *I2F) String() string {
	return "{type：i2f; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<int> 转 %v<double>, 然后放入操作数栈\n", _int, _double)
	stack.PushDouble(_double)
}

func (self *I2D) String() string {
	return "{type：i2d; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<int> 转 %v<long>, 然后放入操作数栈\n", _int, _long)
	stack.PushLong(_long)
}

func (self *I2L) String() string {
	return "{type：i2l; " + self.NoOperandsInstruction.String() + "}"
}
