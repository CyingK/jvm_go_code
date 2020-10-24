package lang

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
	"unsafe"
)

const _constructorConstructorDescriptor = "" +
	"(Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"[Ljava/lang/Class;" +
	"II" +
	"Ljava/lang/String;" +
	"[B[B)V"

func getDeclaredConstructors0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	classObj := vars.GetThis()
	publicOnly := vars.GetBoolean(1)
	class := classObj.GetExtra().(*heap.Class)
	constructors := class.GetConstructors(publicOnly)
	constructorCount := uint(len(constructors))
	classLoader := frame.GetMethod().GetClass().GetClassLoader()
	constructorClass := classLoader.LoadClass("java/lang/reflect/Constructor")
	constructorArr := constructorClass.GetArrayClass().NewArray(constructorCount)
	stack := frame.GetOperandStack()
	stack.PushRef(constructorArr)
	if constructorCount > 0 {
		thread := frame.GetThread()
		constructorObjs := constructorArr.GetRefs()
		constructorInitMethod := constructorClass.GetConstructor(_constructorConstructorDescriptor)
		for i, constructor := range constructors {
			constructorObj := constructorClass.NewObject()
			constructorObj.SetExtra(constructor)
			constructorObjs[i] = constructorObj
			ops := rtda.NewOperandStack(9)
			ops.PushRef(constructorObj)
			ops.PushRef(classObj)
			ops.PushRef(toClassArr(classLoader, constructor.GetParameterTypes()))
			ops.PushRef(toClassArr(classLoader, constructor.GetExceptionTypes()))
			ops.PushInt(int32(constructor.GetAccessFlags()))
			ops.PushInt(int32(0))
			ops.PushRef(getSignatureStr(classLoader, constructor.GetSignature()))
			ops.PushRef(toByteArr(classLoader, constructor.GetAnnotationData()))
			ops.PushRef(toByteArr(classLoader, constructor.GetParameterAnnotationData()))
			shimFrame := rtda.NewShimFrame(thread, ops)
			thread.PushFrame(shimFrame)
			base.InvokeMethod(shimFrame, constructorInitMethod)
		}
	}
}

func toClassArr(loader *heap.ClassLoader, classes []*heap.Class) *heap.Object {
	arrLen := len(classes)
	classArrClass := loader.LoadClass("java/lang/Class").GetArrayClass()
	classArr := classArrClass.NewArray(uint(arrLen))
	if arrLen > 0 {
		classObjs := classArr.GetRefs()
		for i, class := range classes {
			classObjs[i] = class.GetJavaClass()
		}
	}
	return classArr
}

func toByteArr(loader *heap.ClassLoader, goBytes []byte) *heap.Object {
	if goBytes != nil {
		jBytes := castUint8sToInt8s(goBytes)
		return heap.NewByteArray(loader, jBytes)
	}
	return nil
}
func castUint8sToInt8s(goBytes []byte) (jBytes []int8) {
	ptr := unsafe.Pointer(&goBytes)
	jBytes = *((*[]int8)(ptr))
	return
}

func getSignatureStr(loader *heap.ClassLoader, signature string) *heap.Object {
	if signature != "" {
		return heap.ToJavaString(loader, signature)
	}
	return nil
}
