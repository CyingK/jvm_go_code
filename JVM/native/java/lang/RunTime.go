package lang

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"runtime"
)

const JAVA_LANG_RUNTIME = "java/lang/Runtime"

func init() {
	native.Register(JAVA_LANG_RUNTIME, "availableProcessors", "()I", availableProcessors)
}

func availableProcessors(frame *rtda.Frame) {
	numCPU := runtime.NumCPU()
	stack := frame.GetOperandStack()
	stack.PushInt(int32(numCPU))
}
