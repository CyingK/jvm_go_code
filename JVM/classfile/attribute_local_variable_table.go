package classfile

// 局部表量表
type LocalVariableTableEntry struct {
	startPc			uint16		// 起始 PC
	length			uint16		// 长度
	nameIndex		uint16		// 名称索引
	descriptorIndex	uint16		// 描述索引
	index			uint16		// 下标
}

// 属性（局部变量表）
type ATTRIBUTE_LOCAL_VARIABLE_TABLE struct {
	localVariableTable []*LocalVariableTableEntry
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据读入 2 个字节作为局部变量表长度, 创建 []*LocalVariableTableEntry, 对其遍历赋值
func (self *ATTRIBUTE_LOCAL_VARIABLE_TABLE) readInfo(reader *ClassReader) {
	localVariableTableLength := reader.readU2()
	self.localVariableTable = make([]*LocalVariableTableEntry, localVariableTableLength)
	for index := range self.localVariableTable {
		self.localVariableTable[index] = &LocalVariableTableEntry {
			startPc: reader.readU2(),
			length: reader.readU2(),
			nameIndex: reader.readU2(),
			descriptorIndex: reader.readU2(),
			index: reader.readU2(),
		}
	}
}