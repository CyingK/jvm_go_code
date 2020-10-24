package classfile

/*
	CONSTANT_InvokeDynamic_info {
		u1 tag;
		u2 bootstrap_method_attr_index;
		u2 name_and_type_index;
	}
 */

type CONSTANT_InvokeDynamic_info struct {
	bootstrapMethodAttrIndex U2
	nameAndTypeIndex         U2
}

func (self *CONSTANT_InvokeDynamic_info) readInfo(reader *ClassReader) {
	self.bootstrapMethodAttrIndex = reader.readU2()
	self.nameAndTypeIndex = reader.readU2()
}
