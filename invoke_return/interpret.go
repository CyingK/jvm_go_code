package main

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"jvm_go_code/invoke_return/rtda/heap"
)

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	defer catchErr(frame)
	loop(thread, method.Code())
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
	if r := recover(); r != nil {
		fmt.Printf("- LocalVars: %v\n", frame.LocalVars().String())
		fmt.Printf("- OperandStack: %v\n", frame.OperandStack().String())
		panic(r)
	} else {
		fmt.Printf("LocalVars: %v\n", frame.LocalVars().String())
		fmt.Printf("OperandStack: %v\n", frame.OperandStack().String())
	}
}
