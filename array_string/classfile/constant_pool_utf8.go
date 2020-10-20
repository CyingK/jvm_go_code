package classfile

// UTF8 类型常量
type CONSTANT_UTF8_INFO struct {
	str string	// 字符串
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据中读取 2 个字节作为 length, 再从二进制数据中读取 length 个字节，存入 str
func (self *CONSTANT_UTF8_INFO) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// 反编码 UTF8
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
