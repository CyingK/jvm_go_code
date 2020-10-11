package main

import (
	"fmt"
	"jvm_go_code/interpreter/classfile"
	"jvm_go_code/interpreter/instructions"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

func interpret(methodInfo *classfile.MemberInfo) {
	codeAttr := methodInfo.CodeAttribute()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	fmt.Println("Max Locals: ", maxLocals)
	fmt.Println("Max Stack: ", maxStack)
	byteCode := codeAttr.Code()
	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals, maxStack)
	thread.PushFrame(frame)
	defer catchErr(frame)
	loop(thread, byteCode)
}

func loop(thread *rtda.Thread, code []byte) {
	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}
	for {
		println()
		pc := frame.NextPC()
		thread.SetPC(pc)
		reader.Reset(code, pc)
		opCode := reader.ReadUint8()
		instance := instructions.NewInstruction(opCode)
		instance.FetchOperands(reader)
		frame.SetNextPC(reader.PC())
		fmt.Printf("pc: %02d \ninstruction: [0x%02X]%v    ", pc, opCode, instance)
		instance.Execute(frame)
		catchErr(frame)
	}
}

func catchErr(frame *rtda.Frame) {
	fmt.Printf("LocalVars: %v\n", frame.LocalVars().String())
	fmt.Printf("OperandStack: %v\n", frame.OperandStack().String())
	if r := recover(); r != nil {
		panic(r)
	}
}
