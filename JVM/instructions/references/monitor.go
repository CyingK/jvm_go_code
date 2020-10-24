package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type MONITOR_ENTER struct{ base.NoOperandsInstruction }

func (self *MONITOR_ENTER) Execute(frame *rtda.Frame) {
	ref := frame.GetOperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func (self *MONITOR_ENTER) String() string {
	return "{type：monitor_enter; " + self.NoOperandsInstruction.String() + "}\t"
}

type MONITOR_EXIT struct{ base.NoOperandsInstruction }

func (self *MONITOR_EXIT) Execute(frame *rtda.Frame) {
	ref := frame.GetOperandStack().PopRef()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}

func (self *MONITOR_EXIT) String() string {
	return "{type：monitor_exit; " + self.NoOperandsInstruction.String() + "}\t"
}

