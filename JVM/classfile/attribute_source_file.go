package classfile

// 属性（源文件）
type ATTRIBUTE_SOURCE_FILE struct {
	constantPool 	ConstantPool	// 常量池
	sourceFileIndex	uint16			// 源文件索引
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据读入 2 个字节作为源文件索引
func (self *ATTRIBUTE_SOURCE_FILE) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readU2()
}

//--------------------------------------------------------------------Getters

// 获取源文件名
func (self *ATTRIBUTE_SOURCE_FILE) GetFileName() string {
	return self.constantPool.getUtf8(self.sourceFileIndex)
}