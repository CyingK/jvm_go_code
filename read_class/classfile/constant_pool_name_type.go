package classfile

import "log"

/*
 * name_and_type 类型 tag
 */
type CONSTANT_NAME_AND_TYPE_INFO struct {
	nameIndex			uint16	// 名称索引
	descriptorIndex		uint16	// 描述索引
}

/*
 * 从数据中读取 2 次 2 个字节，分别作为 name 和 descriptor 的索引
 */
func (self *CONSTANT_NAME_AND_TYPE_INFO) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
	log.Printf("\t    (4Byte)CONSTANT_name_and_type_info")
	log.Printf("\t\t[name_index: %d], [descriptor_index: %d]", self.nameIndex, self.descriptorIndex)
}
