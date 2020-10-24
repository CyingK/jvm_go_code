package control

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

// void
type RETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	frame.GetThread().PopFrame()
}

func (self *RETURN) String() string {
	return "{type：return; " + self.NoOperandsInstruction.String() + "}\t"
}

// reference
type ARETURN struct {
	base.NoOperandsInstruction
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	ref := currentFrame.GetOperandStack().PopRef()
	invokerFrame.GetOperandStack().PushRef(ref)
}

func (self *ARETURN) String() string {
	return "{type：areturn; " + self.NoOperandsInstruction.String() + "}\t"
}

// double
type DRETURN struct {
	base.NoOperandsInstruction
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	value := currentFrame.GetOperandStack().PopDouble()
	invokerFrame.GetOperandStack().PushDouble(value)
}

func (self *DRETURN) String() string {
	return "{type：dreturn; " + self.NoOperandsInstruction.String() + "}\t"
}

// float
type FRETURN struct {
	base.NoOperandsInstruction
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	value := currentFrame.GetOperandStack().PopFloat()
	invokerFrame.GetOperandStack().PushFloat(value)
}

func (self *FRETURN) String() string {
	return "{type：freturn; " + self.NoOperandsInstruction.String() + "}\t"
}

// int
type IRETURN struct {
	base.NoOperandsInstruction
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	value := currentFrame.GetOperandStack().PopInt()
	invokerFrame.GetOperandStack().PushInt(value)
}


func (self *IRETURN) String() string {
	return "{type：ireturn; " + self.NoOperandsInstruction.String() + "}\t"
}

// long
type LRETURN struct {
	base.NoOperandsInstruction
}

func (self *LRETURN) String() string {
	return "{type：return; " + self.NoOperandsInstruction.String() + "}\t"
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	value := currentFrame.GetOperandStack().PopLong()
	invokerFrame.GetOperandStack().PushLong(value)
}

