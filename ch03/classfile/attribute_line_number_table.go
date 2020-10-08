package classfile

import "log"

type LineNumberTableEntry struct {
	startPc		uint16
	lineNumber	uint16
}

type ATTRIBUTE_LINE_NUMBER_TABLE struct {
	lineNumberTable	[]*LineNumberTableEntry
}

func (self *ATTRIBUTE_LINE_NUMBER_TABLE) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	log.Println("\t\t\t(2Byte)行号表大小：", lineNumberTableLength)
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for index := range self.lineNumberTable {
		log.Printf("\t\t\t    [%d]", index + 1)
		self.lineNumberTable[index] = &LineNumberTableEntry {
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
		log.Println("\t\t\t\t(2Byte)startPC：", self.lineNumberTable[index].startPc)
		log.Println("\t\t\t\t(2Byte)lineNumber：", self.lineNumberTable[index].lineNumber)
	}
}