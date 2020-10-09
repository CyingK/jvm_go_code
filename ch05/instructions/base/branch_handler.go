package base

import "jvm_go_code/ch05/rtda"

func Branch(frame *rtda.Frame, offset int) {
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.Thread().SetPC(nextPC)
}
