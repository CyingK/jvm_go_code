package lang

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

const JAVA_LANG_OBJECT = "java/lang/Object"

func init() {
	native.Register(JAVA_LANG_OBJECT, "getClass", "()Ljava/lang/Class;", getClass)
	native.Register(JAVA_LANG_OBJECT, "hashCode", "()I", hashCode)
	native.Register(JAVA_LANG_OBJECT, "clone", "()Ljava/lang/Object;", clone)
	native.Register(JAVA_LANG_OBJECT, "notifyAll", "()V", notifyAll)
}

func clone(frame *rtda.Frame) {
	this := frame.GetLocalVars().GetThis()
	cloneable := this.GetClass().GetClassLoader().LoadClass("java/lang/Cloneable")
	if !this.GetClass().IsImplements(cloneable) {
		panic("java.lang.CloneNotSupportedException")
	}
	frame.GetOperandStack().PushRef(this.Clone())
}

func hashCode(frame *rtda.Frame) {
	this := frame.GetLocalVars().GetThis()
	class := this.GetClass().GetJavaClass()
	frame.GetOperandStack().PushRef(class)
}

func getClass(frame *rtda.Frame) {
	this := frame.GetLocalVars().GetThis()
	class := this.GetClass().GetJavaClass()
	frame.GetOperandStack().PushRef(class)
}

func notifyAll(frame *rtda.Frame) {
	// todo
}