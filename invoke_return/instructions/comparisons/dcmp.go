package comparisons

// double 比较，因为浮点数计算中可能出现 NaN，固有两个指定来确保程序的正常运行

/********************
 *    dcmpg 		*
 *    dcmpl 		*
 ********************
 * 2				*
 ********************/

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

func util_dprint(v1 float64, v2 float64, result int) {
	fmt.Printf("dcmpx: %v VS %v --> 结果 %d 入操作数栈\n", v1, v2, result)
}

func util_dcmp(frame *rtda.Frame, flag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopDouble()
	v1 := stack.PopDouble()
	if v1 > v2 {
		stack.PushInt(1)
		util_dprint(v1, v2, 1)
	} else if v1 < v2 {
		stack.PushInt(-1)
		util_dprint(v1, v2, -1)
	} else if v1 == v2 {
		stack.PushInt(0)
		util_dprint(v1, v2, 0)
	} else if flag {
		stack.PushInt(1)
		util_dprint(v1, v2, 1)
	} else {
		stack.PushInt(-1)
		util_dprint(v1, v2, -1)
	}
}

type DCMPG struct {
	base.NoOperandsInstruction
}

func (self *DCMPG) String() string {
	return "{type：dcmpg; " + self.NoOperandsInstruction.String() + "}"
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
	return "{type：dcmpg; " + self.NoOperandsInstruction.String() + "}"
}