package classfile

type LineNumberTableEntry struct {
	startPc		uint16
	lineNumber	uint16
}

type ATTRIBUTE_LINE_NUMBER_TABLE struct {
	lineNumberTable	[]*LineNumberTableEntry
}

func (self *ATTRIBUTE_LINE_NUMBER_TABLE) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for index := range self.lineNumberTable {
		self.lineNumberTable[index] = &LineNumberTableEntry {
			startPc: reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}