package classfile

// 名称和描述常量
type CONSTANT_NAME_AND_TYPE_INFO struct {
	nameIndex			uint16	// 名称索引
	descriptorIndex		uint16	// 描述索引
}

//--------------------------------------------------------------------功能类方法

// 从数据中读取 2 次 2 个字节，分别作为 name 和 descriptor 的索引
func (self *CONSTANT_NAME_AND_TYPE_INFO) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
