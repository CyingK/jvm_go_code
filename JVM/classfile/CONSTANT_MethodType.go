package classfile

/*
	CONSTANT_MethodType_info {
		u1 tag;
		u2 descriptor_index;
	}
 */

type CONSTANT_MethodType_info struct {
	descriptorIndex U2
}

func (self *CONSTANT_MethodType_info) readInfo(reader *ClassReader) {
	self.descriptorIndex = reader.readU2()
}
