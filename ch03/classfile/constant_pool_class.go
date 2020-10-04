package classfile

type CONSTANT_CLASS_INFO struct {
	cp			ConstantPool
	nameIndex	uint16
}

func (self *CONSTANT_CLASS_INFO) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

func (self *CONSTANT_CLASS_INFO) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}