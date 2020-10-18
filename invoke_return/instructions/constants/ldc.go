package constants

import (
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

func util_ldc(frame *rtda.Frame, index uint) {
	stack := frame.OperandStack()
	constantPool := frame.Method().Class().ConstantPool()
	constant := constantPool.GetConstant(index)
	switch constant.(type) {
	case int32:
		stack.PushInt(constant.(int32))
	case float32:
		stack.PushFloat(constant.(float32))
	default:
		panic("todo: ldc!")
	}
}

type LDC struct {
	base.Index8Instruction
}

func (self *LDC) Execute(frame *rtda.Frame) {
	util_ldc(frame, self.Index)
}

func (self *LDC) String() string {
	return "{type：ldc; " + self.Index8Instruction.String() + "}"
}

type LDC_W struct {
	base.Index16Instruction
}

func (self *LDC_W) Execute(frame *rtda.Frame) {
	util_ldc(frame, self.Index)
}

func (self *LDC_W) String() string {
	return "{type：ldc_w; " + self.Index16Instruction.String() + "}"
}

type LDC2_W struct {
	base.Index16Instruction
}

func (self *LDC2_W) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	constantPool := frame.Method().Class().ConstantPool()
	constant := constantPool.GetConstant(self.Index)
	switch constant.(type) {
	case int64:
		stack.PushLong(constant.(int64))
	case float64:
		stack.PushDouble(constant.(float64))
	default:
		panic("java.lang.ClassFormatError")
	}
}

func (self *LDC2_W) String() string {
	return "{type：ldc2_w; " + self.Index16Instruction.String() + "}"
}