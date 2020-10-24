package io

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

const JAVA_IO_FILEDESCRIPTOR = "java/io/FileDescriptor"

func init() {
	native.Register(JAVA_IO_FILEDESCRIPTOR, "set", "(I)J", set)
}

func set(frame *rtda.Frame) {
	frame.GetOperandStack().PushLong(0)
}
