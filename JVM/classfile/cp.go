package classfile

/*
 * 生成常量池，并提供相应的 get 方法（package 私有）
 */

// 常量池
type ConstantPool []ConstantInfo

//--------------------------------------------------------------------Getters

// 获取指定下标处的 ConstantInfo
func (self ConstantPool) getConstantInfo(index U2) ConstantInfo {
	if constantInfo := self[index]; constantInfo != nil {
		return constantInfo
	}
	panic("Invalid Constant Pool Index")
}

// 根据 ntInfo 中的 name 索引和 descriptor 索引，在常量池中找到并返回 name 和 type
func (self ConstantPool) getNameAndType(index U2) (string, string) {
	ntInfo := self.getConstantInfo(index).(*CONSTANT_NameAndType_info)
	_name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return _name, _type
}

// 根据 classInfo 中的 name 索引，在常量池中找到并返回 class_name
func (self ConstantPool) getClassName(index U2) string {
	classInfo := self.getConstantInfo(index).(*CONSTANT_Class_info)
	return self.getUtf8(classInfo.nameIndex)
}

// 从 ConstantInfo 中读取 utf8 字符串
func (self ConstantPool) getUtf8(index U2) string {
	utf8Info := self.getConstantInfo(index).(*CONSTANT_Utf8_info)
	return utf8Info.getString()
}


//--------------------------------------------------------------------功能类方法

// 读入常量池的长度, 创建常量信息数组, 对其进行遍历赋值
func resolveConstantPool(reader *ClassReader) ConstantPool {
	constantPoolCount := int(reader.readU2())
	constantPool := make([]ConstantInfo, constantPoolCount)
	for index := 1; index < constantPoolCount; index++ {
		constantPool[index] = readConstantInfo(reader, constantPool)
		switch constantPool[index].(type) {
		case *CONSTANT_Long_info, *CONSTANT_Double_info:
			index++
		}
	}
	return constantPool
}