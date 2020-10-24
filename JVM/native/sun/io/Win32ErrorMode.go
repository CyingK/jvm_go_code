package io

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

func init() {
	native.Register("sun/io/Win32ErrorMode", "setErrorMode", "(J)J", setErrorMode)
}

func setErrorMode(frame *rtda.Frame) {
	frame.GetOperandStack().PushLong(0)
}
