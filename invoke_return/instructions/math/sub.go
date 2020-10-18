package math

// 减法运算

/********************
 *    dsub			*
 *    fsub			*
 *    isub			*
 *    lsub			*
 ********************
 * 4				*
 ********************/

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

/*
 * double 相减
 */
type DSUB struct {
	base.NoOperandsInstruction
}

func (self *DSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 - v2
	fmt.Printf("dsub: 从操作数栈取出 %v, %v, 相减后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushDouble(result)
}

func (self *DSUB) String() string {
	return "{type：dsub; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * float 相减
 */
type FSUB struct {
	base.NoOperandsInstruction
}

func (self *FSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 - v2
	fmt.Printf("fsub: 从操作数栈取出 %v, %v, 相减后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushFloat(result)
}

func (self *FSUB) String() string {
	return "{type：fsub; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * long 相减
 */
type LSUB struct {
	base.NoOperandsInstruction
}

func (self *LSUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 - v2
	fmt.Printf("lsub: 从操作数栈取出 %v, %v, 相减后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushLong(result)
}

func (self *LSUB) String() string {
	return "{type：lsub; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * int 相减
 */
type ISUB struct {
	base.NoOperandsInstruction
}

func (self *ISUB) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 - v2
	fmt.Printf("isub: 从操作数栈取出 %v, %v, 相减后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushInt(result)
}

func (self *ISUB) String() string {
	return "{type：isub; " + self.NoOperandsInstruction.String() + "}"
}