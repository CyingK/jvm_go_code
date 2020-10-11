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
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

/*
 * double 相乘
 */
type DMUL struct {
	base.NoOperandsInstruction
}

func (self *DMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	result := v1 * v2
	fmt.Printf("dmul: 从操作数栈取出 %v, %v, 相乘后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushDouble(result)
}

func (self *DMUL) String() string {
	return "{type：dmul; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * float 相乘
 */
type FMUL struct {
	base.NoOperandsInstruction
}

func (self *FMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	result := v1 * v2
	fmt.Printf("fmul: 从操作数栈取出 %v, %v, 相乘后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushFloat(result)
}

func (self *FMUL) String() string {
	return "{type：fmul; " + self.NoOperandsInstruction.String() + "}"
}


/*
 * long 相乘
 */
type LMUL struct {
	base.NoOperandsInstruction
}

func (self *LMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopLong()
	v1 := stack.PopLong()
	result := v1 * v2
	fmt.Printf("lmul: 从操作数栈取出 %v, %v, 相乘后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushLong(result)
}

func (self *LMUL) String() string {
	return "{type：lmul; " + self.NoOperandsInstruction.String() + "}"
}


/*
 * int 相乘
 */
type IMUL struct {
	base.NoOperandsInstruction
}

func (self *IMUL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	result := v1 * v2
	fmt.Printf("imul: 从操作数栈取出 %v, %v, 相乘后将结果 %v 放入操作数栈\n", v1, v2, result)
	stack.PushInt(result)
}

func (self *IMUL) String() string {
	return "{type：imul; " + self.NoOperandsInstruction.String() + "}"
}
