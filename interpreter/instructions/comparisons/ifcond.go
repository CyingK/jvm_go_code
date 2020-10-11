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
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
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
		fmt.Printf("ifeq: 栈顶标识：%d, 等于 0, 程序计数器跳转至: %d", flag, self.Offset)
		base.Branch(frame, self.Offset)
	}
}

func (self *IFEQ) String() string {
	return "{type：ifeq; " + self.BranchInstruction.String() + "}"
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
		fmt.Printf("ifne: 栈顶标识：%d, 不等于 0, 程序计数器跳转至: %d", flag, self.Offset)
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNE) String() string {
	return "{type：ifne; " + self.BranchInstruction.String() + "}"
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
		fmt.Printf("iflt: 栈顶标识：%d, 小于 0, 程序计数器跳转至: %d", flag, self.Offset)
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLT) String() string {
	return "{type：iflt; " + self.BranchInstruction.String() + "}"
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
		fmt.Printf("ifle: 栈顶标识：%d, 小于等于 0, 程序计数器跳转至: %d", flag, self.Offset)
		base.Branch(frame, self.Offset)
	}
}

func (self *IFLE) String() string {
	return "{type：ifle; " + self.BranchInstruction.String() + "}"
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
		fmt.Printf("ifgt: 栈顶标识：%d, 大于 0, 程序计数器跳转至: %d\n", flag, self.Offset)
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGT) String() string {
	return "{type：ifgt; " + self.BranchInstruction.String() + "}"
}

/*
 * 大于等于
 */
type IFGE struct {
	base.BranchInstruction
}

func (self *IFGE) Execute(frame *rtda.Frame) {
	flag := frame.OperandStack().PopInt()
	if flag >= 0 {
		fmt.Printf("ifge: 栈顶标识：%d, 大于等于 0, 程序计数器跳转至: %d\n", flag, self.Offset)
		base.Branch(frame, self.Offset)
	}
}

func (self *IFGE) String() string {
	return "{type：ifge; " + self.BranchInstruction.String() + "}"
}
