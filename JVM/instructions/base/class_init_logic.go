package base

import (
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

// 初始化类
func InitClass(thread *rtda.Thread, class *heap.Class) {
	class.SetInitStarted()
	scheduleClinit(thread, class)
	initSuperClass(thread, class)
}

// 调用类的 <clinit> 方法
func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		newFrame := thread.NewFrame(clinit)
		thread.PushFrame(newFrame)
	}
}

// 初始化父类：如果 class 不是接口, 则一级一级往上找父类并初始化, 直到 Object 类加载为止
func initSuperClass(thread *rtda.Thread, class *heap.Class) {
	if !class.IsInterface() {
		superClass := class.GetSuperClass()
		if superClass != nil && !superClass.GetInitStarted() {
			InitClass(thread, superClass)
		}
	}
}