package references

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
)

type ANEW_ARRAY struct {
	base.Index16Instruction
}

func (self *ANEW_ARRAY) Execute(frame *rtda.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classRef := constantPool.GetConstant(self.Index).(*heap.ClassRef)
	componentClass := classRef.ResolvedClass()
	stack := frame.GetOperandStack()
	count := stack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}
	arrClass := componentClass.GetArrayClass()
	arr := arrClass.NewArray(uint(count))
	stack.PushRef(arr)
	fmt.Printf("anew_array: 新建 %v 数组, 长度为 %v\n", arrClass.GetName(), count)
}

func (self *ANEW_ARRAY) String() string {
	return "{type：anew_array; " + self.Index16Instruction.String() + "}"
}
