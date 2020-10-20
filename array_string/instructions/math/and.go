package math

// 按位与运算

/********************
 *    iand			*
 *    land			*
 ********************
 * 2				*
 ********************/

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
)

/*
 * int 按位与
 */
type IAND struct {
	base.NoOperandsInstruction
}

func (self *IAND) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 & v2
	fmt.Printf("iand: 从操作数栈取出 %v, %v, 相与后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushInt(result)
}

func (self *IAND) String() string {
	return "{type：iand; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * long 按位与
 */
type LAND struct {
	base.NoOperandsInstruction
}

func (self *LAND) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 & v2
	fmt.Printf("land: 从操作数栈取出 %v, %v, 相与后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushLong(result)
}

func (self *LAND) String() string {
	return "{type：land; " + self.NoOperandsInstruction.String() + "}"
}