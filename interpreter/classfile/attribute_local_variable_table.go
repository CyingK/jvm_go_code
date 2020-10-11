package classfile

type LocalVariableTableEntry struct {
	startPc			uint16
	length			uint16
	nameIndex		uint16
	descriptorIndex	uint16
	index			uint16
}

type ATTRIBUTE_LOCAL_VARIABLE_TABLE struct {
	localVariableTable []*LocalVariableTableEntry
}

func (self *ATTRIBUTE_LOCAL_VARIABLE_TABLE) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readUint16()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for index := range self.localVariableTable {
		self.localVariableTable[index] = &LocalVariableTableEntry {
			startPc: reader.readUint16(),
			length: reader.readUint16(),
			nameIndex: reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index: reader.readUint16(),
		}
	}
}