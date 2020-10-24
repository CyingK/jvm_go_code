package comparisons

// double 比较，因为浮点数计算中可能出现 NaN，固有两个指定来确保程序的正常运行

/********************
 *    dcmpg 		*
 *    dcmpl 		*
 ********************
 * 2				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

func util_dcmp(frame *rtda.Frame, flag bool) {
	stack := frame.GetOperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
	} else if v1 < v2 {
		stack.PushInt(-1)
	} else if v1 == v2 {
		stack.PushInt(0)
	} else if flag {
		stack.PushInt(1)
	} else {
		stack.PushInt(-1)
	}
}

type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) String() string {
	return "{type：dcmpg; " + self.NoOperandsInstruction.String() + "}\t"
}

func (self *DCMPG) Execute(frame *rtda.Frame) {
	util_dcmp(frame, true)
}

type DCMPL struct {
	base.NoOperandsInstruction
}

func (self *DCMPL) Execute(frame *rtda.Frame) {
	util_dcmp(frame, false)
}

func (self *DCMPL) String() string {
	return "{type：dcmpg; " + self.NoOperandsInstruction.String() + "}\t"
}