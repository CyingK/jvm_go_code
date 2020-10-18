package extended

// wide 指令,改变其他指令的行为

/********************
 *    wide			*
 ********************
 * 1				*
 ********************/

import (
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/instructions/loads"
	"jvm_go_code/invoke_return/instructions/math"
	"jvm_go_code/invoke_return/instructions/stores"
	"jvm_go_code/invoke_return/rtda"
)

type WIDE struct {
	modifiedInstruction	base.Instructions
}

func (self WIDE) FetchOperands(reader *base.ByteCodeReader) {
	opcode := reader.ReadUint8()
	switch opcode {
	case 0x15:	// iload
		iload := &loads.ILOAD{}
		iload.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = iload
	case 0x16:	// lload
		lload := &loads.LLOAD{}
		lload.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = lload
	case 0x17:	// fload
		fload := &loads.FLOAD{}
		fload.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = fload
	case 0x18:	// dload
		dload := &loads.DLOAD{}
		dload.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = dload
	case 0x19:	// aload
		aload := &loads.ALOAD{}
		aload.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = aload
	case 0x36:	// istore
		istore := &stores.ISTORE{}
		istore.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = istore
	case 0x37:	// lstore
		lstore := &stores.LSTORE{}
		lstore.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = lstore
	case 0x38:	// fstore
		fstore := &stores.FSTORE{}
		fstore.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = fstore
	case 0x39:	// dstore
		dstore := &stores.DSTORE{}
		dstore.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = dstore
	case 0x3A:	// astore
		astore := &stores.ASTORE{}
		astore.Index = uint(reader.ReadInt16())
		self.modifiedInstruction = astore
	case 0x84:	// iinc
		iinc := &math.IINC{}
		iinc.Index = uint(reader.ReadInt16())
		iinc.Const = int32(reader.ReadInt16())
		self.modifiedInstruction = iinc
	case 0xA9:	// ret
		panic("Unsupported opcode: 0xa9!")
	}
}

func (self *WIDE) Execute(frame *rtda.Frame) {
	self.modifiedInstruction.Execute(frame)
}

func (self *WIDE) String() string {
	return "{type：wide; " + self.modifiedInstruction.String() + "}"
}