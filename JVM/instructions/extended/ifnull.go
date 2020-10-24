package extended

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type IFNULL struct {
	base.BranchInstruction
}

func (self *IFNULL) Execute(frame *rtda.Frame) {
	ref := frame.GetOperandStack().PopRef()
	if ref == nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNULL) String() string {
	if self.Offset < 10 {
		return "{type：ifnull; " + self.BranchInstruction.String() + "}\t\t"
	}
	return "{type：ifnull; " + self.BranchInstruction.String() + "}\t"
}

type IFNONNULL struct {
	base.BranchInstruction
}

func (self *IFNONNULL) Execute(frame *rtda.Frame) {
	ref := frame.GetOperandStack().PopRef()
	if ref != nil {
		base.Branch(frame, self.Offset)
	}
}

func (self *IFNONNULL) String() string {
	return "{type：ifnonnull; " + self.BranchInstruction.String() + "}\t"
}