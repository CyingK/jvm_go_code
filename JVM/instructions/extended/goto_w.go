package extended

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"strconv"
)

type GOTO_W struct {
	offset int
}

func (self *GOTO_W) GetOperands(reader *base.ByteCodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	base.Branch(frame, self.offset)
}

func (self *GOTO_W) String() string {
	return "{type：l2f; offset: " + strconv.Itoa(self.offset) + "}"
}

