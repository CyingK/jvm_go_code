package lang

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

const JAVA_HOME_STRING = "java/lang/String"

func init() {
	native.Register(JAVA_HOME_STRING, "intern", "()Ljava/lang/String;", intern)
}

//
func intern(frame *rtda.Frame) {
	this := frame.GetLocalVars().GetThis()
	interned := heap.InternString(this)
	frame.GetOperandStack().PushRef(interned)
}
