package rtda

import "jvm_go_code/JVM/rtda/heap"

type Frame struct {
	lower        *Frame        // 下一个
	localVars    LocalVars     // 局部变量表
	operandStack *OperandStack // 操作数栈
	thread       *Thread       // 线程
	method       *heap.Method
	nextPC       int            // 程序计数器
}

//--------------------------------------------------------------------构造器
func newFrame(thread *Thread, method *heap.Method,) *Frame {
	return &Frame {
		thread:       thread,
		method:       method,
		localVars:    newLocalVars(method.GetMaxLocals()),
		operandStack: NewOperandStack(method.GetMaxStack()),
	}
}

//--------------------------------------------------------------------Getters
// 获取 localVars
func (self *Frame) GetLocalVars() LocalVars {
	return self.localVars
}

// 获取 operandStack
func (self *Frame) GetOperandStack() *OperandStack {
	return self.operandStack
}


// 获取 thread
func (self *Frame) GetThread() *Thread {
	return self.thread
}

// 获取 method
func (self *Frame) GetMethod() *heap.Method {
	return self.method
}

// 获取 nextPC
func (self *Frame) GetNextPC() int {
	return self.nextPC
}

//--------------------------------------------------------------------Setters
func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}

//--------------------------------------------------------------------功能类方法
func (self *Frame) RevertNextPC() {
	self.nextPC = self.thread.pc
}