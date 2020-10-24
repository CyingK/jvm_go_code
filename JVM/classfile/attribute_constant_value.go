package classfile

// 属性（常量值）
type ATTRIBUTE_CONSTANT_VALUE struct {
	constantValueIndex	uint16		// 常量值索引
}

//--------------------------------------------------------------------Getters

// 获取常量值下标
func (self *ATTRIBUTE_CONSTANT_VALUE) GetConstantValueIndex() uint16 {
	return self.constantValueIndex
}

//--------------------------------------------------------------------功能类方法

// 从 reader 读入 constantValueIndex, 即常量值的下标
func (self *ATTRIBUTE_CONSTANT_VALUE) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readU2()
}