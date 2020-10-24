package rtda

type Stack struct {
	maxSize uint		// 最大栈深
	size    uint		// 已用栈大小
	_top    *Frame		// 栈顶
}

func newStack(maxSize uint) *Stack {
	return &Stack {
		maxSize: maxSize,
	}
}

/*
 * 入栈
 */
func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}

/*
 * 出栈
 */
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("StackEmptyError")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

/*
 * 栈顶
 */
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("StackEmptyError")
	}
	return self._top
}

func (self *Stack) isEmpty() bool {
	return self._top == nil
}

func (self *Stack) getFrames() []*Frame {
	frames := make([]*Frame, 0, self.size)
	for frame := self._top; frame != nil; frame = frame.lower {
		frames = append(frames, frame)
	}
	return frames
}

func (self *Stack) clear() {
	for !self.isEmpty() {
		self.pop()
	}
}
