package comparisons

// 引用的比较与跳转

/********************
 *    if_acmpeq		*
 *    if_acmpne		*
 ********************
 * 2				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

type IF_ACMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ACMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2 {
		fmt.Printf("if_acmpeq: %v == %v = true, 程序计数器跳转至: %v\n", &ref1, &ref2, self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Printf("if_acmpgt: %v == %v = false, 不进行操作\n", &ref1, &ref2)
	}
}

func (self *IF_ACMPEQ) String() string {
	return "{type：if_acmpeq; " + self.BranchInstruction.String() + "}"
}

type IF_ACMPNE struct {
	base.BranchInstruction
}

func (self *IF_ACMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 != ref2 {
		fmt.Printf("if_acmpne: %v != %v = true, 程序计数器跳转至: %v\n", &ref1, &ref2, self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Printf("if_acmpne: %v != %v = false, 不进行操作\n", &ref1, &ref2)
	}
}

func (self *IF_ACMPNE) String() string {
	return "{type：if_acmpne; " + self.BranchInstruction.String() + "}"
}