package classfile

// 属性（未解析的）
type UnparsedAttribute struct {
	name	string	// 属性名
	length	uint32	// 长度
	info	[]byte	// 数据
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据读入 length 个字节
func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}