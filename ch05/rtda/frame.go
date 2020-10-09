package rtda

type Frame struct {
	lower			*Frame			// 下一个
	localVars		LocalVars		// 局部变量表
	operandStack	*OperandStack	// 操作数栈
	thread			*Thread		// 线程
	nextPC			int				// 程序计数器
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func newFrame(thread *Thread, maxLocals uint, maxStack uint) *Frame {
	return &Frame {
		thread: thread,
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) Thread() *Thread {
	return self.thread
}

func (self *Frame) NetPC() int {
	return self.nextPC
}

func (self *Frame) SetNextPC(pc int) {
	self.nextPC = pc
}