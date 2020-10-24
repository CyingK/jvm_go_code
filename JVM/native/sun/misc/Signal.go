package misc

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

func init() {
	_signal(findSignal, "findSignal", "(Ljava/lang/String;)I")
	_signal(handle0, "handle0", "(IJ)J")
}

func _signal(method func(frame *rtda.Frame), name, desc string) {
	native.Register("sun/misc/Signal", name, desc, method)
}

func findSignal(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	vars.GetRef(0) // name
	stack := frame.GetOperandStack()
	stack.PushInt(0) // todo
}

func handle0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	vars.GetInt(0)
	vars.GetLong(1)
	stack := frame.GetOperandStack()
	stack.PushLong(0)
}
