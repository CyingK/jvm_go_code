package base

import (
	"jvm_go_code/ch05/rtda"
)

type Instructions interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
}

/*
 * 无操作指令
 */
type NoOperandsInstruction struct {
	
}

func (self *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {

}

/*
 * 跳转指令
 */
type BranchInstruction struct {
	Offset	int
}

func (self *BranchInstruction) FetchOperands(reader *ByteCodeReader) {
	self.Offset = int(reader.ReadInt16())
}

/*
 * 局部变量表索引
 */
type Index8Instruction struct {
	Index	uint
}

func (self *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint8())
}

type Index16Instruction struct {
	Index	uint
}

func (self *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint16())
}
