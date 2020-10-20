package base

import (
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	println("初始化类：", class.GetName())
	class.SetInitStarted()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.GetSuperClass()
		if superClass != nil && !superClass.GetInitStarted() {
			InitClass(thread, superClass)
		}
	}
}