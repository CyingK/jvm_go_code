package classfile

type CONSTANT_UTF8_INFO struct {
	str string
}

func (self *CONSTANT_UTF8_INFO) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
