package security

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

func init() {
	native.Register("java/security/AccessController", "doPrivileged", "(Ljava/security/PrivilegedAction;)Ljava/lang/Object;", doPrivileged)
	native.Register("java/security/AccessController", "doPrivileged", "(Ljava/security/PrivilegedExceptionAction;)Ljava/lang/Object;", doPrivileged)
	native.Register("java/security/AccessController", "getStackAccessControlContext", "()Ljava/security/AccessControlContext;", getStackAccessControlContext)
}

func doPrivileged(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	action := vars.GetRef(0)
	stack := frame.GetOperandStack()
	stack.PushRef(action)
	method := action.GetClass().GetInstanceMethod("run", "()Ljava/lang/Object;") // todo
	base.InvokeMethod(frame, method)
}

func getStackAccessControlContext(frame *rtda.Frame) {
	frame.GetOperandStack().PushRef(nil)
}
