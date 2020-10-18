package classfile

type ATTRIBUTE_EXCEPTIONS struct {
	exceptionIndexTable	[]uint16
}

func (self *ATTRIBUTE_EXCEPTIONS) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}

func (self *ATTRIBUTE_EXCEPTIONS) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
