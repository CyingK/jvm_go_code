package lang

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"math"
)

const JAVA_LANG_DOUBLE = "java/lang/Double"

func init() {
	native.Register(JAVA_LANG_DOUBLE, "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.Register(JAVA_LANG_DOUBLE, "longBitsToDouble", "(J)D", longBitsToDouble)
}

// double -> long
func doubleToRawLongBits(frame *rtda.Frame) {
	value := frame.GetLocalVars().GetDouble(0)
	bits := math.Float64bits(value) // todo
	frame.GetOperandStack().PushLong(int64(bits))
}

// long -> double
func longBitsToDouble(frame *rtda.Frame) {
	bits := frame.GetLocalVars().GetLong(0)
	value := math.Float64frombits(uint64(bits))
	frame.GetOperandStack().PushDouble(value)
}
