package comparisons

// int 的比较与跳转

/********************
 *    if_icmpeq		*
 *    if_icmpne		*
 *    if_icmplt		*
 *    if_icmple		*
 *    if_icmpgt		*
 *    if_icmpge		*
 ********************
 * 6				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 == v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPEQ) String() string {
	return "{type：if_icmpeq; " + self.BranchInstruction.String() + "}\t"
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPNE) String() string {
	return "{type：if_icmpne; " + self.BranchInstruction.String() + "}\t"
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 < v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLT) String() string {
	return "{type：if_icmplt; " + self.BranchInstruction.String() + "}\t"
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 <= v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPLE) String() string {
	return "{type：if_icmple; " + self.BranchInstruction.String() + "}\t"
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGT) String() string {
	return "{type：if_icmpgt; " + self.BranchInstruction.String() + "}\t"
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 >= v2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ICMPGE) String() string {
	return "{type：if_icmpge; " + self.BranchInstruction.String() + "}\t"
}