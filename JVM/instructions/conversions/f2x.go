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
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * float -> double
 */
type F2D struct {
	base.NoOperandsInstruction
}

func (self *F2D) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	_float := stack.PopFloat()
	_double := float64(_float)
	stack.PushDouble(_double)
}

func (self *F2D) String() string {
	return "{type：f2d; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * float -> int
 */
type F2I struct {
	base.NoOperandsInstruction
}

func (self *F2I) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	_float := stack.PopFloat()
	_int := int32(_float)
	stack.PushInt(_int)
}

func (self *F2I) String() string {
	return "{type：f2i; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * float -> long
 */
type F2L struct {
	base.NoOperandsInstruction
}

func (self *F2L) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	_float := stack.PopFloat()
	_long := int64(_float)
	stack.PushLong(_long)
}

func (self *F2L) String() string {
	return "{type：f2l; " + self.NoOperandsInstruction.String() + "}\t\t"
}