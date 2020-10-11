package constants

// 无操作指令

/********************
 *    nop			*
 ********************
 * 1				*
 ********************/

import (
	"fmt"
	"jvm_go_code/interpreter/instructions/base"
	"jvm_go_code/interpreter/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
	fmt.Println("nop: 啥也没做，哦豁")
}

func (self *NOP) String() string {
	return "{type：nop; nil}"
}