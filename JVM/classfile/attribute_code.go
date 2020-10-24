package classfile

// 属性（代码）
type ATTRIBUTE_CODE struct {
	constantPool   ConstantPool           // 常量池
	maxStack       U2                     // 栈深
	maxLocals      U2                     // 最大局部变量表
	code           []byte                 // 代码
	exceptionTable []*ExceptionTableEntry // 异常表
	attributes     []AttributeInfo        // 属性表
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
	self.maxStack = reader.readU2()
	self.maxLocals = reader.readU2()
	codeLength := reader.readU4()
	self.code = reader.readUn(codeLength)
	self.exceptionTable = getExceptionTable(reader)
	self.attributes = resolveAttributes(reader, self.constantPool)
}

func (self *ATTRIBUTE_CODE) GetLineNumberAttribute() *ATTRIBUTE_LINE_NUMBER_TABLE {
	for _, attrInfo := range self.attributes {
		switch attrInfo.(type) {
		case *ATTRIBUTE_LINE_NUMBER_TABLE:
			return attrInfo.(*ATTRIBUTE_LINE_NUMBER_TABLE)
		}
	}
	return nil
}

// 异常表
type ExceptionTableEntry struct {
	startPc   U2
	endPc     U2
	handlerPc U2
	catchType U2
}

//--------------------------------------------------------------------Getters

// 获取起始 PC
func (self *ExceptionTableEntry) GetStartPC() U2 {
	return self.startPc
}

// 获取末尾 PC
func (self *ExceptionTableEntry) GetEndPC() U2 {
	return self.endPc
}

// 获取 PC 处理器
func (self *ExceptionTableEntry) GetHandlerPC() U2 {
	return self.handlerPc
}

// 获取捕捉类型
func (self *ExceptionTableEntry) GetCatchType() U2 {
	return self.catchType
}