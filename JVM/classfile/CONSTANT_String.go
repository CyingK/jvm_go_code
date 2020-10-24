package classfile

/*
	CONSTANT_String_info {
		u1 tag;
		u2 string_index;
	}
 */

// 字符串类型常量
type CONSTANT_String_info struct {
	constantPool ConstantPool 	// 常量池
	stringIndex  U2				// 字符串索引
}

// 从二进制数据中读取 2 个字节作为 string 索引
func (self *CONSTANT_String_info) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readU2()
}

func (self *CONSTANT_String_info) String() string {
	return self.constantPool.getUtf8(self.stringIndex)
}