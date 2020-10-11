package classfile

type ExceptionTableEntry struct {
	startPc		uint16
	endPc		uint16
	handlerPc	uint16
	catchType	uint16
}

type ATTRIBUTE_CODE struct {
	constantPool			ConstantPool
	maxStack				uint16
	maxLocals				uint16
	code					[]byte
	exceptionTable			[]*ExceptionTableEntry
	attributes				[]AttributeInfo
}

func (self *ATTRIBUTE_CODE) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.constantPool)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTables := make([]*ExceptionTableEntry, exceptionTableLength)
	for index := range exceptionTables {
		exceptionTables[index] = &ExceptionTableEntry {
			startPc: reader.readUint16(),
			endPc: reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
	}
	return exceptionTables
}