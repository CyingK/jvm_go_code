package classfile

/*
	CONSTANT_MethodHandle_info {
		u1 tag;
		u1 reference_kind;
		u2 reference_index;
	}
*/

type CONSTANT_MethodHandle_info struct {
	referenceKind  U1
	referenceIndex U2
}

func (self *CONSTANT_MethodHandle_info) readInfo(reader *ClassReader) {
	self.referenceKind = reader.readU1()
	self.referenceIndex = reader.readU2()
}
