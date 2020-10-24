package math

// 乘法运算

/********************
 *    dmul			*
 *    fmul			*
 *    imul			*
 *    lmul			*
 ********************
 * 4				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * double 相乘
 */
type DMUL struct {
	base.NoOperandsInstruction
}

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	stack.PushDouble(result)
}

func (self *DMUL) String() string {
	return "{type：dmul; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * float 相乘
 */
type FMUL struct {
	base.NoOperandsInstruction
}

func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	stack.PushFloat(result)
}

func (self *FMUL) String() string {
	return "{type：fmul; " + self.NoOperandsInstruction.String() + "}\t\t"
}


/*
 * long 相乘
 */
type LMUL struct {
	base.NoOperandsInstruction
}

func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	stack.PushLong(result)
}

func (self *LMUL) String() string {
	return "{type：lmul; " + self.NoOperandsInstruction.String() + "}\t\t"
}


/*
 * int 相乘
 */
type IMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	stack.PushInt(result)
}

func (self *IMUL) String() string {
	return "{type：imul; " + self.NoOperandsInstruction.String() + "}\t\t"
}
