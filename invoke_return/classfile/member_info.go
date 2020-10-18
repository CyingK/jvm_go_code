package classfile

type MemberInfo struct {
	constantPool    ConstantPool	// 常量池
	accessFlags     uint16			// 访问标识
	nameIndex       uint16			// 名称索引
	descriptorIndex uint16			// 描述索引
	attributes      []AttributeInfo // 属性索引
}

/*
 * 读取所有成员
 */
func readMembers(reader *ClassReader, constantPool ConstantPool) []*MemberInfo {
	memberCount := reader.readUint16()
	members := make([]*MemberInfo, memberCount)
	for index := range members {
		members[index] = readMember(reader, constantPool)
	}
	return members
}

/*
 * 生成一个成员
 */
func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	return &MemberInfo{
		constantPool:    constantPool,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: 	 readAttributes(reader, constantPool),
	}
}

/*
 * 根据 name 索引，在常量池中找到并返回 name
 */
func (self *MemberInfo) Name() string {
	return self.constantPool.getUtf8(self.nameIndex)
}

/*
 * 根据 descriptor 索引，在常量池中找到并返回 descriptor
 */
func (self *MemberInfo) Descriptor() string {
	return self.constantPool.getUtf8(self.descriptorIndex)
}

func (self *MemberInfo) CodeAttribute() *ATTRIBUTE_CODE {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ATTRIBUTE_CODE:
			return attrInfo.(*ATTRIBUTE_CODE)
		}
	}
	return nil
}

func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *MemberInfo) ConstantValueAttribute() *ATTRIBUTE_CONSTANT_VALUE {
	for _, item := range self.attributes {
		switch item.(type) {
		case *ATTRIBUTE_CONSTANT_VALUE:
			return item.(*ATTRIBUTE_CONSTANT_VALUE)
		}
	}
	return nil
}
