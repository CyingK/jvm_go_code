package classfile

type ATTRIBUTE_CONSTANT_VALUE struct {
	constantValueIndex	uint16
}

func (self *ATTRIBUTE_CONSTANT_VALUE) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}

func (self *ATTRIBUTE_CONSTANT_VALUE) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}