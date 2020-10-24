package main

import (
	"fmt"
	"jvm_go_code/JVM/instructions"
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

func interpret(thread *rtda.Thread, logInstance bool) {
	defer func (thread *rtda.Thread) {
		if r := recover(); r != nil {
			logFrames(thread)
			panic(r)
		}
	}(thread)
	loop(thread, logInstance)
}


func loop(thread *rtda.Thread, logInstance bool) {
	reader := &base.ByteCodeReader{}
	for {
		frame := thread.GetCurrentFrame()
		pc := frame.GetNextPC()
		thread.SetPC(pc)
		reader.Reset(frame.GetMethod().GetCode(), pc)
		opCode := reader.ReadUint8()
		instance := instructions.NewInstruction(opCode)
		instance.GetOperands(reader)
		frame.SetNextPC(reader.PC())
		if logInstance {
			logInstruction(frame, instance)
		}
		instance.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
	}
}

func logInstruction(frame *rtda.Frame, instance base.Instructions) {
	method := frame.GetMethod()
	className := method.GetClass().GetName()
	methodName := method.GetName()
	pc := frame.GetThread().GetPC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, instance, instance)
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.GetMethod()
		className := method.GetClass().GetName()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.GetNextPC(), className, method.GetName(), method.GetDescriptor())
	}
}
