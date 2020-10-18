package stack

// 复制栈帧

/********************
 *    dup			*
 *    dup_x1		*
 *    dup_x2		*
 *    dup2			*
 *    dup2_x1		*
 *    dup2_x2		*
 ********************
 * 6				*
 ********************/

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

/*
 * 复制栈顶帧
 */
type DUP struct {
	base.NoOperandsInstruction
}

func (self *DUP) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot := stack.PopSlot()
	stack.PushSlot(slot)
	stack.PushSlot(slot)
	fmt.Printf("dup: 复制栈顶一个元素 %v, [%v] -> [%v, %v]\n",
		slot.String(),
		slot.String(),
		slot.String(), slot.String())
}

func (self *DUP) String() string {
	return "{type：dup; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 复制栈顶帧，插入间隔一个帧的位置
 */
type DUP_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	fmt.Printf("dup_x1: 复制栈顶一个元素 %v, 并隔一个元素存放, [%v, %v] -> [%v, %v, %v]\n",
		slot1.String(),
		slot2.String(), slot1.String(),
		slot1.String(), slot2.String(), slot1.String())
}

func (self *DUP_X1) String() string {
	return "{type：dup_x1; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 复制栈顶帧，插入间隔两个帧的位置
 */
type DUP_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot3)
	fmt.Printf("dup_x2: 复制栈顶一个元素 %v, 并隔两个元素存放, [%v, %v, %v] -> [%v, %v, %v, %v]\n",
		slot1.String(),
		slot3.String(), slot2.String(), slot1.String(),
		slot1.String(), slot3.String(), slot2.String(), slot1.String())
}

func (self *DUP_X2) String() string {
	return "{type：dup_x2; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 复制栈顶两个帧
 */
type DUP2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	fmt.Printf("dup2_x1: 复制栈顶两个元素 %v, %v, [%v, %v] -> [%v, %v, %v, %v]\n",
		slot1.String(), slot2.String(),
		slot2.String(), slot1.String(),
		slot2.String(), slot1.String(), slot2.String(), slot1.String())
}

func (self *DUP2) String() string {
	return "{type：dup2; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 复制栈顶两个帧，插入间隔一个帧的位置
 */
type DUP2_X1 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X1) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	fmt.Printf("dup2_x1: 复制栈顶两个元素 %v, %v, 并隔一个元素存放, [%v, %v, %v] -> [%v, %v, %v, %v, %v]",
		slot1.String(), slot2.String(),
		slot3.String(), slot2.String(), slot1.String(),
		slot2.String(), slot1.String(), slot3.String(), slot2.String(), slot1.String())
}

func (self *DUP2_X1) String() string {
	return "{type：dup2_x1; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 复制栈顶两个帧，插入间隔两个帧的位置
 */
type DUP2_X2 struct {
	base.NoOperandsInstruction
}

func (self *DUP2_X2) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	slot3 := stack.PopSlot()
	slot4 := stack.PopSlot()
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	stack.PushSlot(slot4)
	stack.PushSlot(slot3)
	stack.PushSlot(slot2)
	stack.PushSlot(slot1)
	fmt.Printf("dup2_x2: 复制栈顶两个元素 %v, %v, 并隔两个元素存放, [%v, %v, %v] -> [%v, %v, %v, %v, %v, %v]",
		slot1.String(), slot2.String(),
		slot4.String(), slot3.String(), slot2.String(), slot1.String(),
		slot2.String(), slot1.String(), slot4.String(), slot3.String(), slot2.String(), slot1.String())
}

func (self *DUP2_X2) String() string {
	return "{type：dup2_x2; " + self.NoOperandsInstruction.String() + "}"
}