package comparisons

// 跳转

/********************
 *    ifeq			*
 *    ifne			*
 *    iflt			*
 *    ifle			*
 *    ifgt			*
 *    ifge			*
 ********************
 * 6				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * 等于
 */
type IFEQ struct {
	base.BranchInstruction
}

func (self *IFEQ) Execute(frame *rtda.Frame) {
	flag := frame.GetOperandStack().PopInt()
	if flag == 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFEQ) String() string {
	if self.Offset < 10 {
		return "{type：ifeq; " + self.BranchInstruction.String() + "}\t\t"
	}
	return "{type：ifeq; " + self.BranchInstruction.String() + "}\t"
}

/*
 * 不等于
 */
type IFNE struct {
	base.BranchInstruction
}

func (self *IFNE) Execute(frame *rtda.Frame) {
	flag := frame.GetOperandStack().PopInt()
	if flag != 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNE) String() string {
	if self.Offset < 10 {
		return "{type：ifne; " + self.BranchInstruction.String() + "}\t\t"
	}
	return "{type：ifne; " + self.BranchInstruction.String() + "}\t"
}


/*
 * 小于
 */
type IFLT struct {
	base.BranchInstruction
}

func (self *IFLT) Execute(frame *rtda.Frame) {
	flag := frame.GetOperandStack().PopInt()
	if flag < 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLT) String() string {
	if self.Offset < 10 {
		return "{type：iflt; " + self.BranchInstruction.String() + "}\t\t"
	}
	return "{type：iflt; " + self.BranchInstruction.String() + "}\t"
}


/*
 * 小于等于
 */
type IFLE struct {
	base.BranchInstruction
}

func (self *IFLE) Execute(frame *rtda.Frame) {
	flag := frame.GetOperandStack().PopInt()
	if flag <= 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLE) String() string {
	if self.Offset < 10 {
		return "{type：ifle; " + self.BranchInstruction.String() + "}\t\t"
	}
	return "{type：ifle; " + self.BranchInstruction.String() + "}\t"
}

/*
 * 大于
 */
type IFGT struct {
	base.BranchInstruction
}

func (self *IFGT) Execute(frame *rtda.Frame) {
	flag := frame.GetOperandStack().PopInt()
	if flag > 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGT) String() string {
	if self.Offset < 10 {
		return "{type：ifgt; " + self.BranchInstruction.String() + "}\t\t"
	}
	return "{type：ifgt; " + self.BranchInstruction.String() + "}\t"
}

/*
 * 大于等于
 */
type IFGE struct {
	base.BranchInstruction
}

func (self *IFGE) Execute(frame *rtda.Frame) {
	flag := frame.GetOperandStack().PopInt()
	if flag >= 0 {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGE) String() string {
	if self.Offset < 10 {
		return "{type：ifge; " + self.BranchInstruction.String() + "}\t\t"
	}
	return "{type：ifge; " + self.BranchInstruction.String() + "}\t"
}
