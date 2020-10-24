package math

// 按位或运算

/********************
 *    ior			*
 *    lor			*
 ********************
 * 2				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * int 按位或
 */
type IOR struct {
	base.NoOperandsInstruction
}

func (self *IOR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 | v2
	stack.PushInt(result)
}

func (self *IOR) String() string {
	return "{type：ior; " + self.NoOperandsInstruction.String() + "}\t\t"
}


/*
 * long 按位或
 */
type LOR struct {
	base.NoOperandsInstruction
}

func (self *LOR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 | v2
	stack.PushLong(result)
}

func (self *LOR) String() string {
	return "{type：lor; " + self.NoOperandsInstruction.String() + "}\t\t"
}