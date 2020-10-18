package references

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"jvm_go_code/invoke_return/rtda/heap"
)

type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	methodRef := constantPool.GetConstant(self.Index).(*heap.MethodRef)
	fmt.Println("OK")
	if methodRef.Name() == "println" {
		stack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", stack.PopInt() != 0)
		case "(C)V":
			fmt.Printf("%c\n", stack.PopInt())
		case "(I)V", "(B)V", "(S)V":
			fmt.Printf("%v\n", stack.PopInt())
		case "(F)V":
			fmt.Printf("%v\n", stack.PopFloat())
		case "(J)V":
			fmt.Printf("%v\n", stack.PopLong())
		case "(D)V":
			fmt.Printf("%v\n", stack.PopDouble())
		default:
			panic("println: " + methodRef.Descriptor())
		}
		stack.PopRef()
	}
}

func (self *INVOKE_VIRTUAL) String() string {
	return "{type：invoke_virtual; " + self.Index16Instruction.String() + "}"
}