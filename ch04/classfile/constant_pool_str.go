package classfile

/*
 * string 类型 tag
 */
type CONSTANT_STRING_INFO struct {
	constantPool ConstantPool 	// 常量池
	stringIndex  uint16			// 字符串索引
}

/*
 * 从数据中读取 2 个字节作为 string 索引
 */
func (self *CONSTANT_STRING_INFO) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

/*
 * 根据 string 索引，在常量池中找到并返回 string
 */
func (self *CONSTANT_STRING_INFO) String() string {
	return self.constantPool.getUtf8(self.stringIndex)
}