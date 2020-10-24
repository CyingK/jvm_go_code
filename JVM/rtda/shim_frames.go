package rtda

import "jvm_go_code/JVM/rtda/heap"

func NewShimFrame(thread *Thread, operandStack *OperandStack) *Frame {
	return &Frame {
		thread: 		thread,
		method: 		heap.ShimReturnMethod(),
		operandStack: 	operandStack,
	}
}