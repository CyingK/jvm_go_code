package base

import (
	"jvm_go_code/invoke_return/rtda"
	"strconv"
)

type Instructions interface {
	FetchOperands(reader *ByteCodeReader)
	Execute(frame *rtda.Frame)
	String() string
}

/*
 * 无操作指令
 */
type NoOperandsInstruction struct {
	
}

func (self *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {

}

func (self *NoOperandsInstruction) String() string {
	return "<NPD>"
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

func (self *BranchInstruction) String() string {
	return "BRH<" + strconv.Itoa(self.Offset) + ">"
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

func (self *Index8Instruction) String() string {
	return "Index<" + strconv.Itoa(int(self.Index)) + ">"
}

type Index16Instruction struct {
	Index	uint
}

func (self *Index16Instruction) FetchOperands(reader *ByteCodeReader) {
	self.Index = uint(reader.ReadUint16())
}

func (self *Index16Instruction) String() string {
	return "Index<" + strconv.Itoa(int(self.Index)) + ">"
}