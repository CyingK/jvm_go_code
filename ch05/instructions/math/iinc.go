package math

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
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
	val += self.Const
	localVars.SetInt(self.Index, val)
}