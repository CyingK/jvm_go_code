package comparisons

// 引用的比较与跳转

/********************
 *    if_acmpeq		*
 *    if_acmpne		*
 ********************
 * 2				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPEQ) String() string {
	return "{type：if_acmpeq; " + self.BranchInstruction.String() + "}\t"
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IF_ACMPNE) String() string {
	return "{type：if_acmpne; " + self.BranchInstruction.String() + "}\t"
}