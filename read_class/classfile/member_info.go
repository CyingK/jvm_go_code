package classfile

import "log"

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
	log.Println("(2Byte)读入成员数量：", memberCount)
	members := make([]*MemberInfo, memberCount)
	for index := range members {
		log.Printf("\t[%d]", index + 1)
		members[index] = readMember(reader, constantPool)
	}
	return members
}

/*
 * 生成一个成员
 */
func readMember(reader *ClassReader, constantPool ConstantPool) *MemberInfo {
	memberInfo := &MemberInfo{
		constantPool:    constantPool,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes: 	 nil,
	}
	log.Println("\t    (2Byte)访问标识：", memberInfo.accessFlags)
	log.Println("\t    (2Byte)名称下标：", memberInfo.nameIndex, "//", memberInfo.Name())
	log.Println("\t    (2Byte)描述下标：", memberInfo.descriptorIndex)
	attrbutes := readAttributes(reader, constantPool)
	memberInfo.attributes = attrbutes
	return memberInfo
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