package extended

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.GetOperandStack().PopRef()
	if ref == nil {
		fmt.Println("ifnull: 操作数栈顶为 null, 程序计数器偏移:", self.Offset)
		base.Branch(frame, self.Offset)
	} else {
		fmt.Println("ifnull: 操作数栈顶不为 null, 不进行操作:")
	}
}

func (self *IFNULL) String() string {
	return "{type：ifnull; " + self.BranchInstruction.String() + "}"
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.GetOperandStack().PopRef()
	if ref != nil {
		fmt.Println("ifnonnull: 操作数栈顶不为 null, 程序计数器偏移:", self.Offset)
		base.Branch(frame, self.Offset)
	} else  {
		fmt.Println("ifnonnull: 操作数栈顶为 null, 不进行操作:")
	}
}

func (self *IFNONNULL) String() string {
	return "{type：ifnonnull; " + self.BranchInstruction.String() + "}"
}