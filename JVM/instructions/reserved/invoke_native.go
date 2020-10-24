package reserved

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/native"
	_ "jvm_go_code/JVM/native/java/io"
	_ "jvm_go_code/JVM/native/java/lang"
	_ "jvm_go_code/JVM/native/java/security"
	_ "jvm_go_code/JVM/native/java/util/concurrent/atomic"
	_ "jvm_go_code/JVM/native/sun/io"
	_ "jvm_go_code/JVM/native/sun/misc"
	_ "jvm_go_code/JVM/native/sun/reflect"
	"jvm_go_code/JVM/rtda"
)

// 调用本地方法
type INVOKE_NATIVE struct {
	base.NoOperandsInstruction
}

// 分别取出获取本地方法所需的类名, 方法名, 方法描述, 然后去注册器里找对应方法, 若找到则调用
func (self *INVOKE_NATIVE) Execute(frame *rtda.Frame) {
	method := frame.GetMethod()
	className := method.GetClass().GetName()
	methodName := method.GetName()
	methodDescriptor := method.GetDescriptor()
	methodInfo := className + "." + methodName + ": " + methodDescriptor
	nativeMethod := native.FindNativeMethod(className, methodName, methodDescriptor)
	if nativeMethod == nil {
		panic("java.lang.UnsatisfiedLinkError: " + methodInfo)
	}
	nativeMethod(frame)
}

func (self *INVOKE_NATIVE) String() string {
	return "{type：invoke_native; " + self.NoOperandsInstruction.String() + "}\t"
}