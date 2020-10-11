package classfile

import "log"

type ATTRIBUTE_CONSTANT_VALUE struct {
	constantValueIndex	uint16
}

func (self *ATTRIBUTE_CONSTANT_VALUE) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
	log.Println("\t\t\t\t    (2Byte)常量值下标：", self.constantValueIndex)
}

func (self *ATTRIBUTE_CONSTANT_VALUE) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}