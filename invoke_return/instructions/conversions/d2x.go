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
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
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
	fmt.Printf("%v<double> 转 %v<float>, 然后放入操作数栈\n", _double, _float)
	stack.PushFloat(_float)
}

func (self *D2F) String() string {
	return "{type：d2f; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<double> 转 %v<int>, 然后放入操作数栈\n", _double, _int)
	stack.PushInt(_int)
}

func (self *D2I) String() string {
	return "{type：d2i; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<double> 转 %v<long>, 然后放入操作数栈\n", _double, _long)
	stack.PushLong(_long)
}

func (self *D2L) String() string {
	return "{type：d2l; " + self.NoOperandsInstruction.String() + "}"
}