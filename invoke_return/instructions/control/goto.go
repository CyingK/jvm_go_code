package control

// goto 流程控制

/********************
 *    goto			*
 ********************
 * 1				*
 ********************/

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	fmt.Println("goto: 程序计数器跳转至:", self.Offset)
	base.Branch(frame, self.Offset)
}

func (self *GOTO) String() string {
	return "{type：goto; " + self.BranchInstruction.String() + "}"
}

