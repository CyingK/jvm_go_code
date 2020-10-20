package classfile

// 成员信息
type MemberInfo struct {
	constantPool    ConstantPool		// 常量池
	accessFlags     uint16				// 访问标识
	nameIndex       uint16				// 名称索引
	descriptorIndex uint16				// 描述索引
	attributes      []AttributeInfo 	// 属性索引
}

//--------------------------------------------------------------------Getters

// 根据 name 索引，在常量池中找到并返回 name
func (self *MemberInfo) GetName() string {
	return self.constantPool.getUtf8(self.nameIndex)
}

// 根据 descriptor 索引，在常量池中找到并返回 descriptor
func (self *MemberInfo) GetDescriptor() string {
	return self.constantPool.getUtf8(self.descriptorIndex)
}

// 遍历 attributes, 找到代码属性并返回
func (self *MemberInfo) GetCodeAttribute() *ATTRIBUTE_CODE {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ATTRIBUTE_CODE:
			return attrInfo.(*ATTRIBUTE_CODE)
		}
	}
	return nil
}

// 获取 accessFlags
func (self *MemberInfo) GetAccessFlags() uint16 {
	return self.accessFlags
}

// 遍历 attributes, 找到属性值属性并返回
func (self *MemberInfo) GetConstantValueAttribute() *ATTRIBUTE_CONSTANT_VALUE {
	for _, item := range self.attributes {
		switch item.(type) {
		case *ATTRIBUTE_CONSTANT_VALUE:
			return item.(*ATTRIBUTE_CONSTANT_VALUE)
		}
	}
	return nil
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据读入两个数据作为 member 个数, 创建 MemberInfo 数组, 对其遍历并赋值
func resolveMembers(reader *ClassReader, constantPool ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for index := range members {
		members[index] = readMember(reader, constantPool)
	}
	return members
}

// 生成一个 MemberInfo, 包括 constantPool, accessFlags, nameIndex, descriptorIndex, attributes
func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    constantPool,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      resolveAttributes(reader, constantPool),
	}
}