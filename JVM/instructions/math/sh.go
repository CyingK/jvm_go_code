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
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * int 左移
 */
type ISHL struct {
	base.NoOperandsInstruction
}

func (self *ISHL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1F
	result := v1 << s
	stack.PushInt(result)
}

func (self *ISHL) String() string {
	return "{type：ishl; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * int 算数右移
 */
type ISHR struct {
	base.NoOperandsInstruction
}

func (self *ISHR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1F
	result := v1 >> s
	stack.PushInt(result)
}

func (self *ISHR) String() string {
	return "{type：ishr; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * int 逻辑右移
 */
type IUSHR struct {
	base.NoOperandsInstruction
}

func (self *IUSHR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopInt()
	s := uint32(v2) & 0x1F
	result := int32(uint32(v1) >> s)
	stack.PushInt(result)
}

func (self *IUSHR) String() string {
	return "{type：iushr; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * long 左移
 */
type LSHL struct {
	base.NoOperandsInstruction
}

func (self *LSHL) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3F
	result := v1 << s
	stack.PushLong(result)
}

func (self *LSHL) String() string {
	return "{type：lshl; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * long 算数右移
 */
type LSHR struct {
	base.NoOperandsInstruction
}

func (self *LSHR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3F
	result := v1 >> s
	stack.PushLong(result)
}

func (self *LSHR) String() string {
	return "{type：lshr; " + self.NoOperandsInstruction.String() + "}\t\t"
}

/*
 * long 逻辑右移
 */
type LUSHR struct {
	base.NoOperandsInstruction
}

func (self *LUSHR) Execute(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	v2 := stack.PopInt()
	v1 := stack.PopLong()
	s := uint32(v2) & 0x3F
	result := int64(uint64(v1) >> s)
	stack.PushLong(result)
}

func (self *LUSHR) String() string {
	return "{type：lushr; " + self.NoOperandsInstruction.String() + "}\t\t"
}