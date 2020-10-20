package control

// switch 流程控制,如果case值不可以编码成一个索引表，则实现成lookupswitch指令

/********************
 *    lookup_switch	*
 ********************
 * 1				*
 ********************/

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"strconv"
)

/*
 * 非索引型 switch
 */
type LOOKUP_SWITCH struct {
	defaultOffset	int32
	nparis			int32
	matchOffsets	[]int32
}

func (self *LOOKUP_SWITCH) GetOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.nparis = reader.ReadInt32()
	self.matchOffsets = reader.ReadInt32s(self.nparis * 2)
}

func (self *LOOKUP_SWITCH) Execute(frame *rtda.Frame) {
	key := frame.GetOperandStack().PopInt()
	for i := int32(0); i < self.nparis * 2; i += 2{
		if self.matchOffsets[i] == key {
			offset := self.matchOffsets[i + 1]
			fmt.Println("lookup_switch: 程序计数器跳转至:", offset)
			base.Branch(frame, int(offset))
			return
		}
	}
	base.Branch(frame, int(self.defaultOffset))
}

func (self *LOOKUP_SWITCH) String() string {
	var offsets string
	for index := range self.matchOffsets {
		offsets += strconv.Itoa(int(self.matchOffsets[index]))
	}
	return "{type：lookup_switch; value: " +
		"<default_offset: " + strconv.Itoa(int(self.defaultOffset)) +
		"><nparis: " + strconv.Itoa(int(self.nparis)) +
		"><match_offsets: " + offsets + ">}"
}