package classfile

// 属性（代码）
type ATTRIBUTE_CODE struct {
	constantPool			ConstantPool				// 常量池
	maxStack				uint16						// 栈深
	maxLocals				uint16						// 最大局部变量表
	code					[]byte						// 代码
	exceptionTable			[]*ExceptionTableEntry	// 异常表
	attributes				[]AttributeInfo			// 属性表
}

//--------------------------------------------------------------------Getters
// 获取栈深
func (self *ATTRIBUTE_CODE) GetMaxStack() uint {
	return uint(self.maxStack)
}

// 获取局部变量表长度
func (self *ATTRIBUTE_CODE) GetMaxLocals() uint {
	return uint(self.maxLocals)
}

// 获取代码
func (self *ATTRIBUTE_CODE) GetCode() []byte {
	return self.code
}

// 获取异常表
func (self *ATTRIBUTE_CODE) GetExceptionTable() []*ExceptionTableEntry {
	return self.exceptionTable
}

//--------------------------------------------------------------------功能类方法

// 从 reader 读入 maxStack, maxLocals, code, exeptionTable, attributes
func (self *ATTRIBUTE_CODE) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = getExceptionTable(reader)
	self.attributes = resolveAttributes(reader, self.constantPool)
}

// 异常表
type ExceptionTableEntry struct {
	startPc		uint16
	endPc		uint16
	handlerPc	uint16
	catchType	uint16
}

//--------------------------------------------------------------------Getters

// 获取起始 PC
func (self *ExceptionTableEntry) GetStartPC() uint16 {
	return self.startPc
}

// 获取末尾 PC
func (self *ExceptionTableEntry) GetEndPC() uint16 {
	return self.endPc
}

// 获取 PC 处理器
func (self *ExceptionTableEntry) GetHandlerPC() uint16 {
	return self.handlerPc
}

// 获取捕捉类型
func (self *ExceptionTableEntry) GetCatchType() uint16 {
	return self.catchType
}