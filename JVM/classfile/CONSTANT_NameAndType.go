package classfile

/*
	CONSTANT_NameAndType_info {
		u1 tag;
		u2 name_index;
		u2 descriptor_index;
	}
 */

// 名称和描述常量
type CONSTANT_NameAndType_info struct {
	nameIndex			U2	// 名称索引
	descriptorIndex		U2	// 描述索引
}

// 从数据中读取 2 次 2 个字节，分别作为 name 和 descriptor 的索引
func (self *CONSTANT_NameAndType_info) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readU2()
	self.descriptorIndex = reader.readU2()
}
