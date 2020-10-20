package classfile

// 属性（异常）
type ATTRIBUTE_EXCEPTIONS struct {
	exceptionIndexTable	[]uint16	// 异常表
}

//--------------------------------------------------------------------Getters

// 获取异常下标表
func (self *ATTRIBUTE_EXCEPTIONS) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}

//--------------------------------------------------------------------功能类方法

// 从 reader 读入 exceptionIndexTable 即异常下标表
func (self *ATTRIBUTE_EXCEPTIONS) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}