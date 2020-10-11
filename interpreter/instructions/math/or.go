package math

// 按位或运算

/********************
 *    ior			*
 *    lor			*
 ********************
 * 2				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

/*
 * int 按位或
 */
type IOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	fmt.Printf("ior: 从操作数栈取出 %v , 相或后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushInt(result)
}

func (self *IOR) String() string {
	return "{type：ior; " + self.NoOperandsInstruction.String() + "}"
}


/*
 * long 按位或
 */
type LOR struct {
	base.NoOperandsInstruction
}

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	fmt.Printf("lor: 从操作数栈取出 %v , 相或后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushLong(result)
}

func (self *LOR) String() string {
	return "{type：lor; " + self.NoOperandsInstruction.String() + "}"
}