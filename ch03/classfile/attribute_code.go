package classfile

import "log"

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
	log.Println("\t\t\t    (2Byte)最大栈：", self.maxStack)
	self.maxLocals = reader.readUint16()
	log.Println("\t\t\t    (2Byte)最大局部变量表：", self.maxLocals)
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	log.Printf("\t\t\t    (4Byte)代码长度：(%dByte)", codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.constantPool)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	log.Println("\t\t\t    (2Byte)异常表大小：", exceptionTableLength)
	exceptionTables := make([]*ExceptionTableEntry, exceptionTableLength)
	for index := range exceptionTables {
		exceptionTables[index] = &ExceptionTableEntry {
			startPc: reader.readUint16(),
			endPc: reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}
		log.Printf("\t\t\t\t[%d]", index + 1)
		log.Println("\t\t\t\t    (2Byte)startPC：", exceptionTables[index].startPc)
		log.Println("\t\t\t\t    (2Byte)endPC：", exceptionTables[index].endPc)
		log.Println("\t\t\t\t    (2Byte)handlerPC：", exceptionTables[index].handlerPc)
		log.Println("\t\t\t\t    (2Byte)catchType：", exceptionTables[index].catchType)
	}
	return exceptionTables
}