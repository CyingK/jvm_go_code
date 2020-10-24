package lang

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

const _fieldConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"Ljava/lang/String;" +
	"Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B)V"

func getDeclaredFields0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)
	class := classObj.GetExtra().(*heap.Class)
	fields := class.GetFields(publicOnly)
	fieldCount := uint(len(fields))
	classLoader := frame.GetMethod().GetClass().GetClassLoader()
	fieldClass := classLoader.LoadClass("java/lang/reflect/Field")
	fieldArr := fieldClass.GetArrayClass().NewArray(fieldCount)
	stack := frame.GetOperandStack()
	stack.PushRef(fieldArr)
	if fieldCount > 0 {
		thread := frame.GetThread()
		fieldObjs := fieldArr.GetRefs()
		fieldConstructor := fieldClass.GetConstructor(_fieldConstructorDescriptor)
		for i, goField := range fields {
			fieldObj := fieldClass.NewObject()
			fieldObj.SetExtra(goField)
			fieldObjs[i] = fieldObj
			ops := rtda.NewOperandStack(8)
			ops.PushRef(fieldObj)
			ops.PushRef(classObj)
			ops.PushRef(heap.ToJavaString(classLoader, goField.GetName()))
			ops.PushRef(goField.GetType().GetJavaClass())
			ops.PushInt(int32(goField.GetAccessFlags()))
			ops.PushInt(int32(goField.SlotId()))
			ops.PushRef(getSignatureStr(classLoader, goField.GetSignature()))
			ops.PushRef(toByteArr(classLoader, goField.GetAnnotationData()))
			shimFrame := rtda.NewShimFrame(thread, ops)
			thread.PushFrame(shimFrame)
			base.InvokeMethod(shimFrame, fieldConstructor)
		}
	}
}
