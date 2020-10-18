package classfile

/*
 * member_ref 类型 tag
 */
type CONSTANT_MEMBER_REF_INFO struct {
	constantPool     ConstantPool	// 常量池
	classIndex       uint16			// class索引
	nameAndTypeIndex uint16			// 名称——类型索引
}

/*
 * 从数据中读取 2 次 2 个字节，分别作为 class 和 nameAndType 的索引
 */
func (self *CONSTANT_MEMBER_REF_INFO) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

/*
 * 根据 class 索引，在常量池中找到并返回 class_name
 */
func (self *CONSTANT_MEMBER_REF_INFO) ClassName() string {
	return self.constantPool.getClassName(self.classIndex)
}

/*
 * 根据 nameAndType 索引，在常量池中找到并返回 name_and_descriptor
 */
func (self *CONSTANT_MEMBER_REF_INFO) NameAndDescriptor() (string, string) {
	return self.constantPool.getNameAndType(self.nameAndTypeIndex)
}

/*
 * field 类型 tag
 */
type CONSTANT_FIELD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}

/*
 * method 类型 tag
 */
type CONSTANT_METHOD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}

/*
 * interface_method 类型 tag
 */
type CONSTANT_INTERFACE_METHOD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}