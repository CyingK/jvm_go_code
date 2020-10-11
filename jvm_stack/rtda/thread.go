package rtda

type Thread struct {
	pc		int			// 程序计数器
	stack	*Stack		// 栈
}

func newThread() *Thread {
	return &Thread {
		stack: newStack(1024),
	}
}

/*
 * 获取线程的 PC
 */
func (self *Thread) PC() int {
	return self.pc
}

/*
 * 设置线程的 PC
 */
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

/*
 * 新帧入栈
 */
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

/*
 * 旧帧出栈
 */
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

/*
 * 获取当前栈顶
 */
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}

