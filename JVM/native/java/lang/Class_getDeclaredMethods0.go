package lang

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

const _methodConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"[Ljava/lang/Class;" +
	"Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B[B)V"

func getDeclaredMethods0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)
	class := classObj.GetExtra().(*heap.Class)
	methods := class.GetMethods(publicOnly)
	methodCount := uint(len(methods))
	classLoader := class.GetClassLoader()
	methodClass := classLoader.LoadClass("java/lang/reflect/Method")
	methodArr := methodClass.GetArrayClass().NewArray(methodCount)
	stack := frame.GetOperandStack()
	stack.PushRef(methodArr)
	if methodCount > 0 {
		thread := frame.GetThread()
		methodObjs := methodArr.GetRefs()
		methodConstructor := methodClass.GetConstructor(_methodConstructorDescriptor)
		for i, method := range methods {
			methodObj := methodClass.NewObject()
			methodObj.SetExtra(method)
			methodObjs[i] = methodObj
			ops := rtda.NewOperandStack(12)
			ops.PushRef(methodObj)
			ops.PushRef(classObj)
			ops.PushRef(heap.ToJavaString(classLoader, method.GetName()))
			ops.PushRef(toClassArr(classLoader, method.GetParameterTypes()))
			ops.PushRef(method.GetReturnType().GetJavaClass())
			ops.PushRef(toClassArr(classLoader, method.GetExceptionTypes()))
			ops.PushInt(int32(method.GetAccessFlags()))
			ops.PushInt(int32(0))
			ops.PushRef(getSignatureStr(classLoader, method.GetSignature()))
			ops.PushRef(toByteArr(classLoader, method.GetAnnotationData()))
			ops.PushRef(toByteArr(classLoader, method.GetParameterAnnotationData()))
			ops.PushRef(toByteArr(classLoader, method.GetAnnotationDefaultData()))
			shimFrame := rtda.NewShimFrame(thread, ops)
			thread.PushFrame(shimFrame)
			base.InvokeMethod(shimFrame, methodConstructor)
		}
	}
}
