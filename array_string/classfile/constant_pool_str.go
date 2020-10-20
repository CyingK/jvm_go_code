package classfile

// 字符串类型常量
type CONSTANT_STRING_INFO struct {
	constantPool ConstantPool 	// 常量池
	stringIndex  uint16			// 字符串索引
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据中读取 2 个字节作为 string 索引
func (self *CONSTANT_STRING_INFO) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

//--------------------------------------------------------------------toString

func (self *CONSTANT_STRING_INFO) String() string {
	return self.constantPool.getUtf8(self.stringIndex)
}