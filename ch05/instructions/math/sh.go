package math

// 移位运算

/********************
 *    ishl			*
 *    ishr			*
 *    iushr			*
 *    lshl			*
 *    lshr			*
 *    lushr			*
 ********************
 * 6				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

/*
 * int 左移
 */
type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1F
	result := v1 << s
	stack.PushInt(result)
}

/*
 * int 算数右移
 */
type ISHR struct {
	base.NoOperandsInstruction
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1F
	result := v1 >> s
	stack.PushInt(result)
}

/*
 * int 逻辑右移
 */
type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1F
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

/*
 * long 左移
 */
type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3F
	result := v1 << s
	stack.PushLong(result)
}

/*
 * long 算数右移
 */
type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3F
	result := v1 >> s
	stack.PushLong(result)
}

/*
 * long 逻辑右移
 */
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.OperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3F
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}