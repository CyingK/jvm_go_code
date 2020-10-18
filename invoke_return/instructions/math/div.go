package math

// 除法运算

/********************
 *    ddiv			*
 *    fdiv			*
 *    idiv			*
 *    ldiv			*
 ********************
 * 4				*
 ********************/

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

/*
 * double 相除
 */
type DDIV struct {
	base.NoOperandsInstruction
}

func (self *DDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 / v2
	fmt.Printf("ddiv: 从操作数栈取出 %v, %v, 相除后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushDouble(result)
}

func (self *DDIV) String() string {
	return "{type：ddiv; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * float 相除
 */
type FDIV struct {
	base.NoOperandsInstruction
}

func (self *FDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 / v2
	fmt.Printf("fdiv: 从操作数栈取出 %v, %v, 相除后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushFloat(result)
}

func (self *FDIV) String() string {
	return "{type：ddiv; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * long 相除
 */
type LDIV struct {
	base.NoOperandsInstruction
}

func (self *LDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	fmt.Printf("ldiv: 从操作数栈取出 %v, %v, 相除后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushLong(result)
}

func (self *LDIV) String() string {
	return "{type：ddiv; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * int 相除
 */
type IDIV struct {
	base.NoOperandsInstruction
}

func (self *IDIV) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v2 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}
	result := v1 / v2
	fmt.Printf("idiv: 从操作数栈取出 %v, %v, 相除后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushInt(result)
}

func (self *IDIV) String() string {
	return "{type：ddiv; " + self.NoOperandsInstruction.String() + "}"
}