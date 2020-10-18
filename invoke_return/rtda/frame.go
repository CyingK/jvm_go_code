package rtda

import "jvm_go_code/invoke_return/rtda/heap"

type Frame struct {
	lower        *Frame        // 下一个
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread       // 线程
	method       *heap.Method
	nextPC       int            // 程序计数器
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func newFrame(thread *Thread, method *heap.Method,) *Frame {
	return &Frame {
		thread:       	thread,
		method: 		method,
		localVars:    	newLocalVars(method.MaxLocals()),
		operandStack: 	newOperandStack(method.MaxStack()),
	}
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) Method() *heap.Method {
	return self.method
}

func (self *Frame) NextPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}