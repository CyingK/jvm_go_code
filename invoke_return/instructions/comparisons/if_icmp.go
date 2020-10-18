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
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

type IF_ICMPEQ struct {
	base.BranchInstruction
}

func (self *IF_ICMPEQ) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 == v2 {
		fmt.Printf("if_icmpeq: %d == %d = true, 程序计数器跳转至: %v\n", v1, v2, self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Printf("if_icmpgt: %d == %d = false, 不进行操作\n", v1, v2)
	}
}

func (self *IF_ICMPEQ) String() string {
	return "{type：if_icmpeq; " + self.BranchInstruction.String() + "}"
}

type IF_ICMPNE struct {
	base.BranchInstruction
}

func (self *IF_ICMPNE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 != v2 {
		fmt.Printf("if_icmpne: %d != %d = true, 程序计数器跳转至: %v\n", v1, v2, self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Printf("if_icmpne: %d != %d = false, 不进行操作\n", v1, v2)
	}
}

func (self *IF_ICMPNE) String() string {
	return "{type：if_icmpne; " + self.BranchInstruction.String() + "}"
}

type IF_ICMPLT struct {
	base.BranchInstruction
}

func (self *IF_ICMPLT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 < v2 {
		fmt.Printf("if_icmpeq: %d < %d = true, 程序计数器跳转至: %v\n", v1, v2, self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Printf("if_icmpgt: %d < %d = false, 不进行操作\n", v1, v2)
	}
}

func (self *IF_ICMPLT) String() string {
	return "{type：if_icmplt; " + self.BranchInstruction.String() + "}"
}

type IF_ICMPLE struct {
	base.BranchInstruction
}

func (self *IF_ICMPLE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 <= v2 {
		fmt.Printf("if_icmpeq: %d <= %d = true, 程序计数器跳转至: %v\n", v1, v2, self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Printf("if_icmpgt: %d <= %d = false, 不进行操作\n", v1, v2)
	}
}

func (self *IF_ICMPLE) String() string {
	return "{type：if_icmple; " + self.BranchInstruction.String() + "}"
}

type IF_ICMPGT struct {
	base.BranchInstruction
}

func (self *IF_ICMPGT) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 > v2 {
		fmt.Printf("if_icmpeq: %d > %d = true, 程序计数器跳转至: %v\n", v1, v2, self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Printf("if_icmpgt: %d > %d = false, 不进行操作\n", v1, v2)
	}
}

func (self *IF_ICMPGT) String() string {
	return "{type：if_icmpgt; " + self.BranchInstruction.String() + "}"
}

type IF_ICMPGE struct {
	base.BranchInstruction
}

func (self *IF_ICMPGE) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	if v1 >= v2 {
		fmt.Printf("if_icmpeq: %d >= %d = true, 程序计数器跳转至: %v\n", v1, v2, self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Printf("if_icmpgt: %d >= %d = false, 不进行操作\n", v1, v2)
	}
}

func (self *IF_ICMPGE) String() string {
	return "{type：if_icmpge; " + self.BranchInstruction.String() + "}"
}