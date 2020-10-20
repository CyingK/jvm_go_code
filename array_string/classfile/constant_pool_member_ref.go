package classfile

// 成员引用型常量
type CONSTANT_MEMBER_REF_INFO struct {
	constantPool     ConstantPool	// 常量池
	classIndex       uint16			// class索引
	nameAndTypeIndex uint16			// 名称——类型索引
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据中读取 2 次 2 个字节, 分别作为 class 和 nameAndType 的索引
func (self *CONSTANT_MEMBER_REF_INFO) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

//--------------------------------------------------------------------Getters

// 根据 classIndex, 在常量池中找到并返回 class_name
func (self *CONSTANT_MEMBER_REF_INFO) GetClassName() string {
	return self.constantPool.getClassName(self.classIndex)
}

// 根据 nameAndType 索引, 在常量池中找到 name 和 descriptor 并返回
func (self *CONSTANT_MEMBER_REF_INFO) GetNameAndDescriptor() (string, string) {
	return self.constantPool.getNameAndType(self.nameAndTypeIndex)
}

// 字段类型常量
type CONSTANT_FIELD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}

// 方法类型常量
type CONSTANT_METHOD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}

// 接口方法类型常量
type CONSTANT_INTERFACE_METHOD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}