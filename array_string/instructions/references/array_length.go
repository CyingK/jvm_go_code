package references

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
)

type ARRAY_LENGTH struct {
	base.NoOperandsInstruction
}

func (self *ARRAY_LENGTH) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	arrRef := stack.PopRef()
	if arrRef == nil {
		panic("java.lang.NullPointerException")
	}
	arrLen := arrRef.GetArrayLength()
	stack.PushInt(arrLen)
	fmt.Printf("array_length: 数组类型: %v, 长度: %v\n", arrRef.GetClass().GetName(), arrLen)
}

func (self *ARRAY_LENGTH) String() string {
	return "{type：array_length; " + self.NoOperandsInstruction.String() + "}"
}
