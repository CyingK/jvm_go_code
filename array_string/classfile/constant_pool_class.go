package classfile

// 类类型常量
type CONSTANT_CLASS_INFO struct {
	constantPool ConstantPool	// 常量池
	nameIndex    uint16			// 名称索引
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据中读取 2 个字节，作为名称索引
func (self *CONSTANT_CLASS_INFO) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}

//--------------------------------------------------------------------Getters

// 根据 name 索引，在常量池中找到并返回 name
func (self *CONSTANT_CLASS_INFO) GetName() string {
	return self.constantPool.getUtf8(self.nameIndex)
}