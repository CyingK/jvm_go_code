package rtda

type Frame struct {
	lower			*Frame			// 下一个
	localVars		LocalVars		// 局部变量表
	operandStack	*OperandStack	// 操作数栈
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}

func NewFrame(maxLocals uint, maxStack uint) *Frame {
	return &Frame {
		localVars: newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

