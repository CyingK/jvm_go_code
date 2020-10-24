package classfile

/*
	CONSTANT_Utf8_info {
		u1 tag;
		u2 length;
		u1 bytes[length];
	}
 */

// UTF8 类型常量
type CONSTANT_Utf8_info struct {
	length		U2
	bytes		[]byte
}

// 从二进制数据中读取 2 个字节作为 length, 再从二进制数据中读取 length 个字节，存入 str
func (self *CONSTANT_Utf8_info) readInfo(reader *ClassReader) {
	self.length = reader.readU2()
	self.bytes = reader.readUn(U4(self.length))
}

// 反编码 UTF8
func (self *CONSTANT_Utf8_info) getString() string {
	return string(self.bytes)
}
