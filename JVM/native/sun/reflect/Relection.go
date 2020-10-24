package reflect

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

func init() {
	native.Register("sun/reflect/Reflection", "getCallerClass", "()Ljava/lang/Class;", getCallerClass)
	native.Register("sun/reflect/Reflection", "getClassAccessFlags", "(Ljava/lang/Class;)I", getClassAccessFlags)

}

func getCallerClass(frame *rtda.Frame) {
	callerFrame := frame.GetThread().GetFrames()[2] // todo
	callerClass := callerFrame.GetMethod().GetClass().GetJavaClass()
	frame.GetOperandStack().PushRef(callerClass)
}

func getClassAccessFlags(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	_type := vars.GetRef(0)
	goClass := _type.GetExtra().(*heap.Class)
	flags := goClass.GetAccessFlags()
	stack := frame.GetOperandStack()
	stack.PushInt(int32(flags))
}
