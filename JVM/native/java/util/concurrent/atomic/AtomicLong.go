package atomic

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

func init() {
	native.Register("java/util/concurrent/atomic/AtomicLong", "VMSupportsCS8", "()Z", vmSupportsCS8)
}

func vmSupportsCS8(frame *rtda.Frame) {
	frame.GetOperandStack().PushBoolean(false)
}
