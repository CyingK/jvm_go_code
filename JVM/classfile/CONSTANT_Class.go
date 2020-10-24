package classfile

/*
	CONSTANT_Class_info {
		U1 tag;
		U2 name_index;
	}
 */

// 类类型常量
type CONSTANT_Class_info struct {
	tag				U1					// 类型标记
	nameIndex    	U2           		// 名称索引
	constantPool 	ConstantPool 		// 常量池, 便于 GetName()
}

// 从二进制数据中读取 2 个字节，作为名称索引
func (self *CONSTANT_Class_info) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readU2()
}

// 根据 name 索引，在常量池中找到并返回 name
func (self *CONSTANT_Class_info) GetName() string {
	return self.constantPool.getUtf8(self.nameIndex)
}