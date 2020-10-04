package classfile

type CONSTANT_MEMBER_REF_INFO struct {
	cp					ConstantPool
	classIndex			uint16
	nameAndTypeIndex	uint16
}

func (self *CONSTANT_MEMBER_REF_INFO) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}

func (self *CONSTANT_MEMBER_REF_INFO) ClassName() string {
	return self.cp.getUtf8(self.classIndex)
}

func (self *CONSTANT_MEMBER_REF_INFO) NameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)
}

type CONSTANT_FIELD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}

type CONSTANT_METHOD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}

type CONSTANT_INTERFACE_METHOD_REF_INFO struct {
	CONSTANT_MEMBER_REF_INFO
}