package classfile

type CONSTANT_STRING_INFO struct {
	cp 			ConstantPool
	stringIndex	uint16
}

func (self *CONSTANT_STRING_INFO) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

func (self *CONSTANT_STRING_INFO) String() string {
	return self.cp.getUtf8(self.stringIndex)
}