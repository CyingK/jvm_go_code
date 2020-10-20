package base

import (
	"fmt"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
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
	if method.IsNative() {
		if method.GetName() == "registerNatives" {
			thread.PopFrame()
		} else {
			panic(fmt.Sprintf("native method: %v.%v%v\n",
				method.GetClass().GetName(), method.GetName(), method.Descriptor()))
		}
	}
}