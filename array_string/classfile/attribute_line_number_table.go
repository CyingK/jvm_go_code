package classfile

// 行号表
type LineNumberTableEntry struct {
	startPc		uint16		// 起始 PC
	lineNumber	uint16		// 行号
}

// 属性（行号表）
type ATTRIBUTE_LINE_NUMBER_TABLE struct {
	lineNumberTable	[]*LineNumberTableEntry
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据读入 2 个字节作为行号表长度, 创建 []*LineNumberTableEntry, 对其遍历赋值
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