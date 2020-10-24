package lang

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

func init() {
	native.Register("java/lang/Thread", "currentThread", "()Ljava/lang/Thread;", currentThread)
	native.Register("java/lang/Thread", "setPriority0", "(I)V", setPriority0)
	native.Register("java/lang/Thread", "isAlive", "()Z", isAlive)
	native.Register("java/lang/Thread", "start0", "()V", start0)
}

func currentThread(frame *rtda.Frame) {
	classLoader := frame.GetMethod().GetClass().GetClassLoader()
	threadClass := classLoader.LoadClass("java/lang/Thread")
	jThread := threadClass.NewObject()
	threadGroupClass := classLoader.LoadClass("java/lang/ThreadGroup")
	jGroup := threadGroupClass.NewObject()
	jThread.SetRefVar("group", "Ljava/lang/ThreadGroup;", jGroup)
	jThread.SetIntVar("priority", "I", 1)
	frame.GetOperandStack().PushRef(jThread)
}

func setPriority0(frame *rtda.Frame) {
}

func isAlive(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	stack.PushBoolean(false)
}

func start0(frame *rtda.Frame) {
}
