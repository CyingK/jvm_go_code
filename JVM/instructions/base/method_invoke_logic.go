package base

import (
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.GetThread()
	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)
	argSlotCount := int(method.GetArgSlotCount())
	if argSlotCount > 0 {
		for item := argSlotCount - 1; item >= 0; item-- {
			slot := invokerFrame.GetOperandStack().PopSlot()
			newFrame.GetLocalVars().SetSlot(uint(item), slot)
		}
	}
}