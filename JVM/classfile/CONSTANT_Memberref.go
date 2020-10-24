package classfile

// 成员引用型常量
type CONSTANT_Memberref_info struct {
	constantPool     ConstantPool	// 常量池
	classIndex       U2				// class索引
	nameAndTypeIndex U2				// 名称——类型索引
}

// 从二进制数据中读取 2 次 2 个字节, 分别作为 class 和 nameAndType 的索引
func (self *CONSTANT_Memberref_info) readInfo(reader *ClassReader) {
	self.classIndex = reader.readU2()
	self.nameAndTypeIndex = reader.readU2()
}

// 根据 classIndex, 在常量池中找到并返回 class_name
func (self *CONSTANT_Memberref_info) GetClassName() string {
	return self.constantPool.getClassName(self.classIndex)
}

// 根据 nameAndType 索引, 在常量池中找到 name 和 descriptor 并返回
func (self *CONSTANT_Memberref_info) GetNameAndDescriptor() (string, string) {
	return self.constantPool.getNameAndType(self.nameAndTypeIndex)
}