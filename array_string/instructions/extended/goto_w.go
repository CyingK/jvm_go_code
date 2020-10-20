package extended

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"strconv"
)

type GOTO_W struct {
	offset int
}

func (self *GOTO_W) GetOperands(reader *base.ByteCodeReader) {
	self.offset = int(reader.ReadInt32())
}

func (self *GOTO_W) Execute(frame *rtda.Frame) {
	fmt.Println("goto_w: 程序计数器偏移:", self.offset)
	base.Branch(frame, self.offset)
}

func (self *GOTO_W) String() string {
	return "{type：l2f; offset: " + strconv.Itoa(self.offset) + "}"
}
