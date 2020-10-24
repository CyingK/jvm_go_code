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
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * double 取反
 */
type DNEG struct {
	base.NoOperandsInstruction
}

func (self *DNEG) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopDouble()
	stack.PushDouble(-value)
}

func (self *DNEG) String() string {
	return "{type：dneg; " + self.NoOperandsInstruction.String() + "}\t\t"
}


/*
 * float 取反
 */
type FNEG struct {
	base.NoOperandsInstruction
}

func (self *FNEG) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopFloat()
	stack.PushFloat(-value)
}

func (self *FNEG) String() string {
	return "{type：fneg; " + self.NoOperandsInstruction.String() + "}\t\t"
}


/*
 * long 取反
 */
type LNEG struct {
	base.NoOperandsInstruction
}

func (self *LNEG) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopLong()
	stack.PushLong(-value)
}

func (self *LNEG) String() string {
	return "{type：lneg; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * int 取反
 */
type INEG struct {
	base.NoOperandsInstruction
}

func (self *INEG) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	value := stack.PopInt()
	stack.PushInt(-value)
}

func (self *INEG) String() string {
	return "{type：ineg; " + self.NoOperandsInstruction.String() + "}\t\t"
}
