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
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
)

/*
 * long -> float
 */
type L2F struct {
	base.NoOperandsInstruction
}

func (self *L2F) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	_long := stack.PopLong()
	_float := float32(_long)
	fmt.Printf("%v<long> 转 %v<float>, 然后放入操作数栈\n", _long, _float)
	stack.PushFloat(_float)
}

func (self *L2F) String() string {
	return "{type：l2f; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * long -> int
 */
type L2I struct {
	base.NoOperandsInstruction
}

func (self *L2I) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	_long := stack.PopLong()
	_int := int32(_long)
	fmt.Printf("%v<long> 转 %v<int>, 然后放入操作数栈\n", _long, _int)
	stack.PushInt(_int)
}

func (self *L2I) String() string {
	return "{type：l2i; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * long -> double
 */
type L2D struct {
	base.NoOperandsInstruction
}

func (self *L2D) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	_long := stack.PopLong()
	_double := float64(_long)
	fmt.Printf("%v<long> 转 %v<double>, 然后放入操作数栈\n", _long, _double)
	stack.PushDouble(_double)
}

func (self *L2D) String() string {
	return "{type：l2d; " + self.NoOperandsInstruction.String() + "}"
}