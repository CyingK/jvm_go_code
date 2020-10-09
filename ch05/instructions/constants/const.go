package constants

// const 指令，把隐含在操作码中的常量值推入操作数栈顶

/********************
 *    aconst_null	*
 *    dconst_0		*
 *    dconst_1		*
 *    fconst_0		*
 *    fconst_1		*
 *    fconst_2		*
 *    iconst_0		*
 *    iconst_1		*
 *    iconst_2		*
 *    iconst_3		*
 *    iconst_4		*
 *    iconst_5		*
 *    iconst_m1		*
 *    lconst_0		*
 *    lconst_1		*
 ********************
 * 15				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

/*
 * null -> 操作数栈顶
 */
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

/*
 * 0.0 -> 操作数栈顶
 */
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

/*
 * 1.0 -> 操作数栈顶
 */
type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

/*
 * 0.0f -> 操作数栈顶
 */
type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
}

/*
 * 1.0f -> 操作数栈顶
 */
type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
}

/*
 * 2.0f -> 操作数栈顶
 */
type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
}

/*
 * -1 -> 操作数栈顶
 */
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

/*
 * 0 -> 操作数栈顶
 */
type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

/*
 * 1 -> 操作数栈顶
 */
type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}

/*
 * 2 -> 操作数栈顶
 */
type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

/*
 * 3 -> 操作数栈顶
 */
type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

/*
 * 4 -> 操作数栈顶
 */
type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

/*
 * 5 -> 操作数栈顶
 */
type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

/*
 * 0L -> 操作数栈顶
 */
type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

/*
 * 1L -> 操作数栈顶
 */
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}
