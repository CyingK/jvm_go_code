package control

// switch 流程控制,如果case值可以编码成一个索引表，则实现成tableswitch指令

/********************
 *    table_switch	*
 ********************
 * 1				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"strconv"
)

/*
 * 索引型 switch
 */
type TABLE_SWITCH struct {
	defaultOffset		int32
	low					int32
	high				int32
	jumpOffset			[]int32
}

func (self *TABLE_SWITCH) GetOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()
	jumpOffsetCount := self.high - self.low
	self.jumpOffset = reader.ReadInt32s(jumpOffsetCount)
}

func (self *TABLE_SWITCH) Execute(frame *rtda.Frame) {
	index := frame.GetOperandStack().PopInt()
	var offset int
	if index >= self.low && index <= self.high {
		offset = int(self.jumpOffset[index - self.low])
	} else {
		offset = int(self.defaultOffset)
	}
	base.Branch(frame, offset)
}

func (self *TABLE_SWITCH) String() string {
	var offsets string
	for index := range self.jumpOffset {
		offsets += strconv.Itoa(int(self.jumpOffset[index]))
	}
	return "{type：table_switch; value: " +
		"<default_offset: " + strconv.Itoa(int(self.defaultOffset)) +
		"><low: " + strconv.Itoa(int(self.low)) +
		"><high: " + strconv.Itoa(int(self.high)) +
		"><match_offsets: " + offsets + ">}"
}