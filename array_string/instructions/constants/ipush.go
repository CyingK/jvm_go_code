package constants

//bipush 指令从操作数中获取一个 byte 型整数，扩展成 int 型，然后推入栈顶
//sipush 指令从操作数中获取一个 short 型整数，扩展成 int 型，然后推入栈顶。

/********************
 *    bipush		*
 *    sipush		*
 ********************
 * 2				*
 ********************/

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"strconv"
)

/*
 * 操作数 -> 操作数栈顶
 */
type BIPUSH struct {
	value		int8
}

func (self *BIPUSH) GetOperands(reader *base.ByteCodeReader) {
	self.value = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame) {
	value := int32(self.value)
	fmt.Printf("bipush: %v 入操作数栈\n", value)
	frame.GetOperandStack().PushInt(value)
}

func (self *BIPUSH) String() string {
	return "{type：bipush; value: " + strconv.Itoa(int(self.value)) + "}"
}

/*
 * 操作数 -> 操作数栈顶
 */
type SIPUSH struct {
	value		int16
}

func (self *SIPUSH) GetOperands(reader *base.ByteCodeReader) {
	self.value = reader.ReadInt16()
}

func (self *SIPUSH) Execute(frame *rtda.Frame) {
	value := int32(self.value)
	fmt.Printf("sipush: %v 入操作数栈\n", value)
	frame.GetOperandStack().PushInt(value)
}

func (self *SIPUSH) String() string {
	return "{type：sipush; value: " + strconv.Itoa(int(self.value)) + "}"
}