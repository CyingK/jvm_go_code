package constants

// 无操作指令

/********************
 *    nop			*
 ********************
 * 1				*
 ********************/

import (
	"jvm_go_code/ch05/instructions/base"
	"jvm_go_code/ch05/rtda"
)

type NOP struct {
	base.NoOperandsInstruction
}

func (self *NOP) Excute(frame *rtda.Frame) {

}
