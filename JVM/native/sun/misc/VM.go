package misc

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

func init() {
	native.Register("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	classLoader := frame.GetMethod().GetClass().GetClassLoader()
	javaLangSystemClass := classLoader.LoadClass("java/lang/System")
	initSystemClass := javaLangSystemClass.GetStaticMethod("initializeSystemClass", "()V")
	base.InvokeMethod(frame, initSystemClass)
}