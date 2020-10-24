package constants

// 无操作指令

/********************
 *    nop			*
 ********************
 * 1				*
 ********************/

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame) {
}

func (self *NOP) String() string {
	return "{type：nop; nil}"
}