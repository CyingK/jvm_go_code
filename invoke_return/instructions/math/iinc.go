package math

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"strconv"
)

// 给局部变量表中的int变量增加常量值，局部变量表索引和常量值都由指令的操作数提供

/********************
 *    iinc			*
 ********************
 * 1				*
 ********************/

/*
 * 加上一个 int 数
 */
type IINC struct {
	Index 	uint
	Const	int32
}

func (self *IINC) FetchOperands(reader *base.ByteCodeReader) {
	self.Index = uint(reader.ReadInt8())
	self.Const = int32(reader.ReadInt8())
}

func (self *IINC) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	fmt.Printf("iinc: 从操作数栈取出 %v 加 %d 变为 %v 放入操作数栈\n", val, self.Const, val + self.Const)
	val += self.Const
	localVars.SetInt(self.Index, val)
}

func (self *IINC) String() string {
	return "{type：iinc; " +
		"Index: " + strconv.Itoa(int(self.Index)) +
		", Const: " + strconv.Itoa(int(self.Const)) +
		"}"
}