package reflect

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

func init() {
	native.Register("sun/reflect/NativeConstructorAccessorImpl", "newInstance0", "(Ljava/lang/reflect/Constructor;[Ljava/lang/Object;)Ljava/lang/Object;", newInstance0)
}

func newInstance0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	constructorObj := vars.GetRef(0)
	argArrObj := vars.GetRef(1)
	goConstructor := getGoConstructor(constructorObj)
	goClass := goConstructor.GetClass()
	if !goClass.GetInitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.GetThread(), goClass)
		return
	}
	obj := goClass.NewObject()
	stack := frame.GetOperandStack()
	stack.PushRef(obj)
	ops := convertArgs(obj, argArrObj, goConstructor)
	shimFrame := rtda.NewShimFrame(frame.GetThread(), ops)
	frame.GetThread().PushFrame(shimFrame)
	base.InvokeMethod(shimFrame, goConstructor)
}

func getGoMethod(methodObj *heap.Object) *heap.Method {
	return _getGoMethod(methodObj, false)
}

func getGoConstructor(constructorObj *heap.Object) *heap.Method {
	return _getGoMethod(constructorObj, true)
}

func _getGoMethod(methodObj *heap.Object, isConstructor bool) *heap.Method {
	extra := methodObj.GetExtra()
	if extra != nil {
		return extra.(*heap.Method)
	}
	if isConstructor {
		root := methodObj.GetRefVar("root", "Ljava/lang/reflect/Constructor;")
		return root.GetExtra().(*heap.Method)
	} else {
		root := methodObj.GetRefVar("root", "Ljava/lang/reflect/Method;")
		return root.GetExtra().(*heap.Method)
	}
}

func convertArgs(this, argArr *heap.Object, method *heap.Method) *rtda.OperandStack {
	if method.GetArgSlotCount() == 0 {
		return nil
	}
	ops := rtda.NewOperandStack(method.GetArgSlotCount())
	if !method.IsStatic() {
		ops.PushRef(this)
	}
	if method.GetArgSlotCount() == 1 && !method.IsStatic() {
		return ops
	}
	return ops
}
