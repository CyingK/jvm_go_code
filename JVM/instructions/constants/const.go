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
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

/*
 * null -> 操作数栈顶
 */
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushRef(nil)
}

func (self *ACONST_NULL) String() string {
	return "{type：aconst_null; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 0.0 -> 操作数栈顶
 */
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushDouble(0.0)
}

func (self *DCONST_0) String() string {
	return "{type：dconst_0; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 1.0 -> 操作数栈顶
 */
type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushDouble(1.0)
}

func (self *DCONST_1) String() string {
	return "{type：dconst_1; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 0.0f -> 操作数栈顶
 */
type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushFloat(0.0)
}

func (self *FCONST_0) String() string {
	return "{type：fconst_0; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 1.0f -> 操作数栈顶
 */
type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushFloat(1.0)
}

func (self *FCONST_1) String() string {
	return "{type：fconst_1; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 2.0f -> 操作数栈顶
 */
type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushFloat(2.0)
}

func (self *FCONST_2) String() string {
	return "{type：fconst_2; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * -1 -> 操作数栈顶
 */
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(-1)
}

func (self *ICONST_M1) String() string {
	return "{type：iconst_m1; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 0 -> 操作数栈顶
 */
type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(0)
}

func (self *ICONST_0) String() string {
	return "{type：iconst_0; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 1 -> 操作数栈顶
 */
type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(1)
}

func (self *ICONST_1) String() string {
	return "{type：iconst_1; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 2 -> 操作数栈顶
 */
type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(2)
}

func (self *ICONST_2) String() string {
	return "{type：iconst_2; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 3 -> 操作数栈顶
 */
type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(3)
}

func (self *ICONST_3) String() string {
	return "{type：iconst_3; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 4 -> 操作数栈顶
 */
type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(4)
}

func (self *ICONST_4) String() string {
	return "{type：iconst_4; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 5 -> 操作数栈顶
 */
type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushInt(5)
}

func (self *ICONST_5) String() string {
	return "{type：iconst_5; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 0L -> 操作数栈顶
 */
type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushLong(0)
}

func (self *LCONST_0) String() string {
	return "{type：lconst_0; " + self.NoOperandsInstruction.String() + "}\t"
}

/*
 * 1L -> 操作数栈顶
 */
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.GetOperandStack().PushLong(1)
}

func (self *LCONST_1) String() string {
	return "{type：lconst_1; " + self.NoOperandsInstruction.String() + "}\t"
}