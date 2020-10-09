package classfile

/*
 * 生成常量池，并提供相应的 get 方法（package 私有）
 */

type ConstantPool []ConstantInfo

/*
 * 读取常量池
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	constantPoolCount := int(reader.readUint16())
	constantPool := make([]ConstantInfo, constantPoolCount)
	for index := 1; index < constantPoolCount; index++ {
		constantPool[index] = readConstantInfo(reader, constantPool)
		switch constantPool[index].(type) {
		case *CONSTANT_LONG_INFO, *CONSTANT_DOUBLE_INFO:
			index++
		}
	}
	return constantPool
}

/*
 * 获取指定下标处的 ConstantInfo
 */
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if constantInfo := self[index]; constantInfo != nil {
		return constantInfo
	}
	panic("Invalid Constant Pool Index")
}

/*
 * 根据 ntInfo 中的 name 索引和 descriptor 索引，在常量池中找到并返回 name 和 type
 */
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*CONSTANT_NAME_AND_TYPE_INFO)
	_name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return _name, _type
}

/*
 * 根据 classInfo 中的 name 索引，在常量池中找到并返回 class_name
 */
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*CONSTANT_CLASS_INFO)
	return self.getUtf8(classInfo.nameIndex)
}

/*
 * 从 ConstantInfo 中读取 utf8 字符串
 */
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*CONSTANT_UTF8_INFO)
	return utf8Info.str
}