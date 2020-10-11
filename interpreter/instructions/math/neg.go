package math

// 取反运算

/********************
 *    dneg			*
 *    fneg			*
 *    ineg			*
 *    lneg			*
 ********************
 * 4				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

/*
 * double 取反
 */
type DNEG struct {
	base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopDouble()
	fmt.Printf("dneg: 从操作数栈取出 %v , 取反后将结果 %v 放入操作数栈\n", value, -value)
	stack.PushDouble(-value)
}

func (self *DNEG) String() string {
	return "{type：dneg; " + self.NoOperandsInstruction.String() + "}"
}


/*
 * float 取反
 */
type FNEG struct {
	base.NoOperandsInstruction
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopFloat()
	fmt.Printf("fneg: 从操作数栈取出 %v , 取反后将结果 %v 放入操作数栈\n", value, -value)
	stack.PushFloat(-value)
}

func (self *FNEG) String() string {
	return "{type：fneg; " + self.NoOperandsInstruction.String() + "}"
}


/*
 * long 取反
 */
type LNEG struct {
	base.NoOperandsInstruction
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopLong()
	fmt.Printf("lneg: 从操作数栈取出 %v , 取反后将结果 %v 放入操作数栈\n", value, -value)
	stack.PushLong(-value)
}

func (self *LNEG) String() string {
	return "{type：lneg; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * int 取反
 */
type INEG struct {
	base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	value := stack.PopInt()
	fmt.Printf("ineg: 从操作数栈取出 %v , 取反后将结果 %v 放入操作数栈\n", value, -value)
	stack.PushInt(-value)
}

func (self *INEG) String() string {
	return "{type：ineg; " + self.NoOperandsInstruction.String() + "}"
}
