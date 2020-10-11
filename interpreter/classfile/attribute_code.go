package classfile

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

func (self *ATTRIBUTE_CODE) MaxStack() uint {
	return uint(self.maxStack)
}

func (self *ATTRIBUTE_CODE) MaxLocals() uint {
	return uint(self.maxLocals)
}

func (self *ATTRIBUTE_CODE) Code() []byte {
	return self.code
}

func (self *ATTRIBUTE_CODE) ExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

type ExceptionTableEntry struct {
	startPc		uint16
	endPc		uint16
	handlerPc	uint16
	catchType	uint16
}

func (self *ExceptionTableEntry) StartPC() uint16 {
	return self.startPc
}

func (self *ExceptionTableEntry) EndPC() uint16 {
	return self.endPc
}

func (self *ExceptionTableEntry) HandlerPC() uint16 {
	return self.handlerPc
}

func (self *ExceptionTableEntry) CatchType() uint16 {
	return self.catchType
}