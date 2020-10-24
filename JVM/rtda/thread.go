package rtda

import "jvm_go_code/JVM/rtda/heap"

type Thread struct {
	pc		int			// 程序计数器
	stack	*Stack		// 栈
}

// 清空当前线程的所有栈帧
func (self *Thread) ClearStack() {
	self.stack.clear()
}

// 获取当前栈顶的栈帧
func (self *Thread) GetCurrentFrame() *Frame {
	return self.stack.top()
}

// 获取当前线程的所有栈帧
func (self *Thread) GetFrames() []*Frame {
	return self.stack.getFrames()
}

// 获取 pc
func (self *Thread) GetPC() int {
	return self.pc
}

// 获取 stack
func (self *Thread) GetStack() *Stack {
	return self.stack
}

// 获取当前栈顶的栈帧
func (self *Thread) GetTopFrame() *Frame {
	return self.stack.top()
}

// 当前线程的虚拟机栈为空: true, 反之: false
func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}

// 创建新的栈帧
func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

// 从虚拟机栈弹出一个栈帧
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// 将栈帧推入当前线程的虚拟机栈
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// 设置 pc
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

// 创建新的线程, 默认虚拟机栈最多存储 1024 个栈帧
func NewThread() *Thread {
	return &Thread {
		stack: newStack(1024),
	}
}