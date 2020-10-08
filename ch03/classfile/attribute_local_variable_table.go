package classfile

import "log"

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
	log.Println("\t\t\t(2Byte)局部变量表大小：", localVariableTableLength)
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for index := range self.localVariableTable {
		self.localVariableTable[index] = &LocalVariableTableEntry {
			startPc: reader.readUint16(),
			length: reader.readUint16(),
			nameIndex: reader.readUint16(),
			descriptorIndex: reader.readUint16(),
			index: reader.readUint16(),
		}
		log.Printf("\t\t\t    [%d]", index + 1)
		log.Println("\t\t\t\t(2Byte)startPC：", self.localVariableTable[index].startPc)
		log.Println("\t\t\t\t(2Byte)length：", self.localVariableTable[index].length)
		log.Println("\t\t\t\t(2Byte)nameIndex：", self.localVariableTable[index].nameIndex)
		log.Println("\t\t\t\t(2Byte)descriptorIndex：", self.localVariableTable[index].descriptorIndex)
		log.Println("\t\t\t\t(2Byte)index：", self.localVariableTable[index].index)
	}
}