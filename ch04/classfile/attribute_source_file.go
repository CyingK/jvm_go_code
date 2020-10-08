package classfile

type ATTRIBUTE_SOURCE_FILE struct {
	constantPool 	ConstantPool
	sourceFileIndex	uint16
}

func (self *ATTRIBUTE_SOURCE_FILE) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *ATTRIBUTE_SOURCE_FILE) FileName() string {
	return self.constantPool.getUtf8(self.sourceFileIndex)
}