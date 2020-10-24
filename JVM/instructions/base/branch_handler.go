package base

import "jvm_go_code/JVM/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.GetThread().GetPC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
