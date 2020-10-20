package comparisons

// float 比较，因为浮点数计算中可能出现 NaN，固有两个指定来确保程序的正常运行

/********************
 *    fcmpg			*
 *    fcmpl			*
 ********************
 * 2				*
 ********************/

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
)

func util_fprint(v1 float32, v2 float32, result int) {
	fmt.Printf("fcmpx: %v VS %v --> 结果 %d 入操作数栈\n", v1, v2, result)
}

func util_fcmp(frame *rtda.Frame, flag bool) {
	stack := frame.GetOperandStack()
	v2 := stack.PopFloat()
	v1 := stack.PopFloat()
	if v1 > v2 {
		stack.PushInt(1)
		util_fprint(v1, v2, 1)
	} else if v1 < v2 {
		stack.PushInt(-1)
		util_fprint(v1, v2, -1)
	} else if v1 == v2 {
		stack.PushInt(0)
		util_fprint(v1, v2, 0)
	} else if flag {
		stack.PushInt(1)
		util_fprint(v1, v2, 1)
	} else {
		stack.PushInt(-1)
		util_fprint(v1, v2, -1)
	}
}

type FCMPG struct {
	base.NoOperandsInstruction
}

func (self *FCMPG) Execute(frame *rtda.Frame) {
	util_fcmp(frame, true)
}

func (self *FCMPG) String() string {
	return "{type：fcmpg; " + self.NoOperandsInstruction.String() + "}"
}

type FCMPL struct {
	base.NoOperandsInstruction
}

func (self *FCMPL) Execute(frame *rtda.Frame) {
	util_fcmp(frame, false)
}

func (self *FCMPL) String() string {
	return "{type：fcmpl; " + self.NoOperandsInstruction.String() + "}"
}