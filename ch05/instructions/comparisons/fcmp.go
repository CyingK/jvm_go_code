package comparisons

// float 比较，因为浮点数计算中可能出现 NaN，固有两个指定来确保程序的正常运行

/********************
 *    fcmpg			*
 *    fcmpl			*
 ********************
 * 2				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

func util_fcmp(frame *rtda.Frame, flag bool) {
	stack := frame.OperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
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

type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	util_fcmp(frame, true)
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	util_fcmp(frame, false)
}