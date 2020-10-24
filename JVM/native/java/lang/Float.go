package lang

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"math"
)

const JAVA_LANG_FLOAT = "java/lang/Float"

func init() {
	native.Register(JAVA_LANG_FLOAT, "floatToRawIntBits", "(F)I", floatToRawIntBits)
	native.Register(JAVA_LANG_FLOAT, "intBitsToFloat", "(I)F", intBitsToFloat)
}

// int -> float
func intBitsToFloat(frame *rtda.Frame) {
	bits := frame.GetLocalVars().GetInt(0)
	value := math.Float32frombits(uint32(bits))
	frame.GetOperandStack().PushFloat(value)
}

// float -> int
func floatToRawIntBits(frame *rtda.Frame) {
	value := frame.GetLocalVars().GetFloat(0)
	bits := math.Float32bits(value)
	frame.GetOperandStack().PushInt(int32(bits))
}
