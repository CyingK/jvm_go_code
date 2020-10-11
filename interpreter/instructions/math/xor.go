package math

// 按位异或运算

/********************
 *    ixor			*
 *    lxor			*
 ********************
 * 2				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

/*
 * int 按位异或
 */
type IXOR struct {
	base.NoOperandsInstruction
}

func (self *IXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 ^ v2
	fmt.Printf("ixor: 从操作数栈取出 %v, %v, 相异或后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushInt(result)
}

func (self *IXOR) String() string {
	return "{type：ixor; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * long 按位异或
 */
type LXOR struct {
	base.NoOperandsInstruction
}

func (self *LXOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 ^ v2
	fmt.Printf("lxor: 从操作数栈取出 %v, %v, 相异或后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushLong(result)
}

func (self *LXOR) String() string {
	return "{type：lxor; " + self.NoOperandsInstruction.String() + "}"
}