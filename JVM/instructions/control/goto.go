package control

// goto 流程控制

/********************
 *    goto			*
 ********************
 * 1				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type GOTO struct {
	base.BranchInstruction
}

func (self *GOTO) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.Offset)
}

func (self *GOTO) String() string {
	if self.Offset < 10 {
		return "{type：goto; " + self.BranchInstruction.String() + "}\t\t"
	}
	return "{type：goto; " + self.BranchInstruction.String() + "}\t"
}

