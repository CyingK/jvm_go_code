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
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
)

/*
 * null -> 操作数栈顶
 */
type ACONST_NULL struct {
	base.NoOperandsInstruction
}

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
	fmt.Println("aconst_null: 把 null 放入操作数栈")
}

func (self *ACONST_NULL) String() string {
	return "{type：aconst_null; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 0.0 -> 操作数栈顶
 */
type DCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
	fmt.Println("dconst_0: 把 0.0 放入操作数栈")
}

func (self *DCONST_0) String() string {
	return "{type：dconst_0; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 1.0 -> 操作数栈顶
 */
type DCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *DCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
	fmt.Println("dconst_1: 把 1.0 放入操作数栈")
}

func (self *DCONST_1) String() string {
	return "{type：dconst_1; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 0.0f -> 操作数栈顶
 */
type FCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0.0)
	fmt.Println("fconst_0: 把 0.0f 放入操作数栈")
}

func (self *FCONST_0) String() string {
	return "{type：fconst_0; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 1.0f -> 操作数栈顶
 */
type FCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1.0)
	fmt.Println("fconst_1: 把 1.0f 放入操作数栈")
}

func (self *FCONST_1) String() string {
	return "{type：fconst_1; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 2.0f -> 操作数栈顶
 */
type FCONST_2 struct {
	base.NoOperandsInstruction
}

func (self *FCONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2.0)
	fmt.Println("fconst_2: 把 2.0f 放入操作数栈")
}

func (self *FCONST_2) String() string {
	return "{type：fconst_2; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * -1 -> 操作数栈顶
 */
type ICONST_M1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
	fmt.Println("iconst_m1: 把 -1 放入操作数栈")
}

func (self *ICONST_M1) String() string {
	return "{type：iconst_m1; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 0 -> 操作数栈顶
 */
type ICONST_0 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
	fmt.Println("iconst_0: 把 0 放入操作数栈")
}

func (self *ICONST_0) String() string {
	return "{type：iconst_0; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 1 -> 操作数栈顶
 */
type ICONST_1 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
	fmt.Println("iconst_1: 把 1 放入操作数栈")
}

func (self *ICONST_1) String() string {
	return "{type：iconst_1; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 2 -> 操作数栈顶
 */
type ICONST_2 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
	fmt.Println("iconst_2: 把 2 放入操作数栈")
}

func (self *ICONST_2) String() string {
	return "{type：iconst_2; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 3 -> 操作数栈顶
 */
type ICONST_3 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
	fmt.Println("iconst_3: 把 3 放入操作数栈")
}

func (self *ICONST_3) String() string {
	return "{type：iconst_3; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 4 -> 操作数栈顶
 */
type ICONST_4 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
	fmt.Println("iconst_4: 把 4 放入操作数栈")
}

func (self *ICONST_4) String() string {
	return "{type：iconst_4; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 5 -> 操作数栈顶
 */
type ICONST_5 struct {
	base.NoOperandsInstruction
}

func (self *ICONST_5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
	fmt.Println("iconst_5: 把 5 放入操作数栈")
}

func (self *ICONST_5) String() string {
	return "{type：iconst_5; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 0L -> 操作数栈顶
 */
type LCONST_0 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
	fmt.Println("lconst_0: 把 0L 放入操作数栈")
}

func (self *LCONST_0) String() string {
	return "{type：lconst_0; " + self.NoOperandsInstruction.String() + "}"
}

/*
 * 1L -> 操作数栈顶
 */
type LCONST_1 struct {
	base.NoOperandsInstruction
}

func (self *LCONST_1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
	fmt.Println("lconst_1: 把 1L 放入操作数栈")
}

func (self *LCONST_1) String() string {
	return "{type：lconst_1; " + self.NoOperandsInstruction.String() + "}"
}