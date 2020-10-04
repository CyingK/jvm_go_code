package classfile

type CONSTANT_NAME_AND_TYPE_INFO struct {
	nameIndex			uint16
	descriptorIndex		uint16
}

func (self *CONSTANT_NAME_AND_TYPE_INFO) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
