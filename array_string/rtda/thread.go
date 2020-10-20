package rtda

import "jvm_go_code/array_string/rtda/heap"

type Thread struct {
	pc		int			// 程序计数器
	stack	*Stack		// 栈
}

//--------------------------------------------------------------------构造器
func NewThread() *Thread {
	return &Thread {
		stack: newStack(1024),
	}
}

func (self *Thread) NewFrame(method *heap.Method) *Frame {
	return newFrame(self, method)
}

//--------------------------------------------------------------------Getters
// 获取 pc
func (self *Thread) GetPC() int {
	return self.pc
}

// 获取 stack
func (self *Thread) GetStack() *Stack {
	return self.stack
}

//--------------------------------------------------------------------Setters
// 设置 pc
func (self *Thread) SetPC(pc int) {
	self.pc = pc
}

//--------------------------------------------------------------------判断类方法
// 判断当前栈是否为空
func (self *Thread) IsStackEmpty() bool {
	return self.stack.isEmpty()
}

//--------------------------------------------------------------------功能类方法
// 新帧入栈
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

// 旧帧出栈
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

// 获取当前栈顶
func (self *Thread) GetCurrentFrame() *Frame {
	return self.stack.top()
}

// 获取当前栈顶
func (self *Thread) GetTopFrame() *Frame {
	return self.stack.top()
}