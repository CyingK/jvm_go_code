package math

// 求余运算

/********************
 *    drem			*
 *    frem			*
 *    irem			*
 *    lrem			*
 ********************
 * 4				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"math"
)

/*
 * double 取余
 */
type DREM struct {
	base.NoOperandsInstruction
}

func (self *DREM) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := math.Mod(v1, v2)
	stack.PushDouble(result)
}

func (self *DREM) String() string {
	return "{type：drem; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * float 取余
 */
type FREM struct {
	base.NoOperandsInstruction
}

func (self *FREM) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := float32(math.Mod(float64(v1), float64(v2)))
	stack.PushFloat(result)
}

func (self *FREM) String() string {
	return "{type：frem; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * int 取余
 */
type IREM struct {
	base.NoOperandsInstruction
}

func (self *IREM) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushInt(result)
}

func (self *IREM) String() string {
	return "{type：irem; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * long 取余
 */
type LREM struct {
	base.NoOperandsInstruction
}

func (self *LREM) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 % v2
	stack.PushLong(result)
}

func (self *LREM) String() string {
	return "{type：lrem; " + self.NoOperandsInstruction.String() + "}\t\t"
}