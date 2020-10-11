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
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
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
	fmt.Printf("%v<float> 转 %v<double>, 然后放入操作数栈\n", _float, _double)
	stack.PushDouble(_double)
}

func (self *F2D) String() string {
	return "{type：f2d; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<float> 转 %v<int>, 然后放入操作数栈\n", _float, _int)
	stack.PushInt(_int)
}

func (self *F2I) String() string {
	return "{type：f2i; " + self.NoOperandsInstruction.String() + "}"
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
	fmt.Printf("%v<float> 转 %v<long>, 然后放入操作数栈\n", _float, _long)
	stack.PushLong(_long)
}

func (self *F2L) String() string {
	return "{type：f2l; " + self.NoOperandsInstruction.String() + "}"
}