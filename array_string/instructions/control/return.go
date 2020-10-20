package control

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
)

// void
type RETURN struct {
	base.NoOperandsInstruction
}

func (self *RETURN) Execute(frame *rtda.Frame) {
	fmt.Println("return: 当前方法的栈帧出栈")
	frame.GetThread().PopFrame()
}

func (self *RETURN) String() string {
	return "{type：return; " + self.NoOperandsInstruction.String() + "}"
}

// reference
type ARETURN struct {
	base.NoOperandsInstruction
}

func (self *ARETURN) Execute(frame *rtda.Frame) {
	fmt.Println("areturn: 弹出当前栈帧, 获取调用者栈帧, 将当前栈帧操作数栈的引用推入调用者栈帧")
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	ref := currentFrame.GetOperandStack().PopRef()
	invokerFrame.GetOperandStack().PushRef(ref)
}

func (self *ARETURN) String() string {
	return "{type：areturn; " + self.NoOperandsInstruction.String() + "}"
}

// double
type DRETURN struct {
	base.NoOperandsInstruction
}

func (self *DRETURN) Execute(frame *rtda.Frame) {
	fmt.Println("dreturn: 弹出当前栈帧, 获取调用者栈帧, 将当前栈帧操作数栈的 double 值推入调用者栈帧")
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	value := currentFrame.GetOperandStack().PopDouble()
	invokerFrame.GetOperandStack().PushDouble(value)
}

func (self *DRETURN) String() string {
	return "{type：dreturn; " + self.NoOperandsInstruction.String() + "}"
}

// float
type FRETURN struct {
	base.NoOperandsInstruction
}

func (self *FRETURN) Execute(frame *rtda.Frame) {
	fmt.Println("freturn: 弹出当前栈帧, 获取调用者栈帧, 将当前栈帧操作数栈的 float 值推入调用者栈帧")
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	value := currentFrame.GetOperandStack().PopFloat()
	invokerFrame.GetOperandStack().PushFloat(value)
}

func (self *FRETURN) String() string {
	return "{type：freturn; " + self.NoOperandsInstruction.String() + "}"
}

// int
type IRETURN struct {
	base.NoOperandsInstruction
}

func (self *IRETURN) Execute(frame *rtda.Frame) {
	fmt.Println("ireturn: 弹出当前栈帧, 获取调用者栈帧, 将当前栈帧操作数栈的 int 值推入调用者栈帧")
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	value := currentFrame.GetOperandStack().PopInt()
	invokerFrame.GetOperandStack().PushInt(value)
}


func (self *IRETURN) String() string {
	return "{type：ireturn; " + self.NoOperandsInstruction.String() + "}"
}

// long
type LRETURN struct {
	base.NoOperandsInstruction
}

func (self *LRETURN) String() string {
	return "{type：return; " + self.NoOperandsInstruction.String() + "}"
}

func (self *LRETURN) Execute(frame *rtda.Frame) {
	fmt.Println("lreturn: 弹出当前栈帧, 获取调用者栈帧, 将当前栈帧操作数栈的 long 值推入调用者栈帧")
	thread := frame.GetThread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.GetTopFrame()
	value := currentFrame.GetOperandStack().PopLong()
	invokerFrame.GetOperandStack().PushLong(value)
}

