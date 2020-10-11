package math

// 加法运算

/********************
 *    dadd			*
 *    fadd			*
 *    iadd			*
 *    ladd			*
 ********************
 * 4				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

/*
 * double 相加
 */
type DADD struct {
	base.NoOperandsInstruction
}

func (self *DADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 + v2
	fmt.Printf("dadd: 从操作数栈取出 %v, %v, 相加后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushDouble(result)
}

func (self *DADD) String() string {
	return "{type：dadd; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * float 相加
 */
type FADD struct {
	base.NoOperandsInstruction
}

func (self *FADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 + v2
	fmt.Printf("fadd: 从操作数栈取出 %v, %v, 相加后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushFloat(result)
}

func (self *FADD) String() string {
	return "{type：fadd; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * long 相加
 */
type LADD struct {
	base.NoOperandsInstruction
}

func (self *LADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 + v2
	fmt.Printf("ladd: 从操作数栈取出 %v, %v, 相加后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushLong(result)
}

func (self *LADD) String() string {
	return "{type：ladd; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * int 相加
 */
type IADD struct {
	base.NoOperandsInstruction
}

func (self *IADD) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 + v2
	fmt.Printf("iadd: 从操作数栈取出 %v, %v, 相加后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushInt(result)
}

func (self *IADD) String() string {
	return "{type：ladd; " + self.NoOperandsInstruction.String() + "}"
}