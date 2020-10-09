package comparisons

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

/*
 * 等于
 */
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	flag := frame.OperandStack().PopInt()
	if flag == 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
 * 不等于
 */
type IFNE struct {
	base.BranchInstruction
}

func (self *IFNE) Execute(frame *rtda.Frame) {
	flag := frame.OperandStack().PopInt()
	if flag != 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
 * 小于
 */
type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	flag := frame.OperandStack().PopInt()
	if flag < 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
 * 小于等于
 */
type IFLE struct {
	base.BranchInstruction
}

func (self *IFLE) Execute(frame *rtda.Frame) {
	flag := frame.OperandStack().PopInt()
	if flag <= 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
 * 大于
 */
type IFGT struct {
	base.BranchInstruction
}

func (self *IFGT) Execute(frame *rtda.Frame) {
	flag := frame.OperandStack().PopInt()
	if flag > 0 {
		base.Branch(frame, self.Offset)
	}
}

/*
 * 大于等于
 */
type IFGE struct {
	base.BranchInstruction
}

func (self *IFGE) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}