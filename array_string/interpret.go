package main

import (
	"fmt"
	"jvm_go_code/array_string/instructions"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
)

func interpret(method *heap.Method, logInstance bool, args []string) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)
	java_args := createArgsArray(method.GetClass().GetClassLoader(), args)
	frame.GetLocalVars().SetRef(0, java_args)
	defer catchErr(thread)
	loop(thread, logInstance)
}

func createArgsArray(classLoader *heap.ClassLoader, args []string) *heap.Object {
	stringClass := classLoader.LoadClass("java/lang/String")
	argsArray := stringClass.GetArrayClass().NewArray(uint(len(args)))
	java_args := argsArray.GetRefs()
	for index, item := range args {
		java_args[index] = heap.JString(classLoader, item)
	}
	return argsArray
}

func loop(thread *rtda.Thread, logInstance bool) {
	reader := &base.ByteCodeReader{}
	var count = 1
	for {
		count++
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
		fmt.Printf("Round: %d\tPc: %02d\tInstruction: [0x%02X]%v  ", count, pc, opCode, instance)
		instance.Execute(frame)
		if thread.IsStackEmpty() {
			break
		}
		catchErr(thread)
	}
}

func logInstruction(frame *rtda.Frame, instance base.Instructions) {
	method := frame.GetMethod()
	className := method.GetClass().GetName()
	methodName := method.GetName()
	pc := frame.GetThread().GetPC()
	fmt.Printf("%v.%v() #%2d %T %v\n", className, methodName, pc, instance, instance)
}

func catchErr(thread *rtda.Thread) {
	if r := recover(); r != nil {
		logFrames(thread)
		panic(r)
	}
}

func logFrames(thread *rtda.Thread) {
	for !thread.IsStackEmpty() {
		frame := thread.PopFrame()
		method := frame.GetMethod()
		className := method.GetClass().GetName()
		fmt.Printf(">> pc:%4d %v.%v%v \n", frame.GetNextPC(), className, method.GetName(), method.Descriptor())
	}
}
