package classfile

type CONSTANT_METHOD_HANDLE_INFO struct {
	referenceKind  uint8
	referenceIndex uint16
}

func (self *CONSTANT_METHOD_HANDLE_INFO) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readUint8()
	self.referenceIndex = reader.readUint16()
}

type CONSTANT_METHOD_TYPE_INFO struct {
	descriptorIndex uint16
}

func (self *CONSTANT_METHOD_TYPE_INFO) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readUint16()
}

type CONSTANT_INVOKE_DYNAMIC_INFO struct {
	bootstrapMethodAttrIndex uint16
	nameAndTypeIndex         uint16
}

func (self *CONSTANT_INVOKE_DYNAMIC_INFO) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
