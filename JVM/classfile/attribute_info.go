package classfile

// 属性接口
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

//--------------------------------------------------------------------构造器

// 根据属性名创建不同的属性并返回
func newAttributeInfo(attrName string, attrLength uint32, constantPool ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &ATTRIBUTE_CODE{
			constantPool: constantPool,
		}
	case "ConstantValue":
		return &ATTRIBUTE_CONSTANT_VALUE{}
	case "Deprecated":
		return &ATTRIBUTE_DEPRECATED{}
	case "Exceptions":
		return &ATTRIBUTE_EXCEPTIONS{}
	case "LineNumberTable":
		return &ATTRIBUTE_LINE_NUMBER_TABLE{}
	case "LocalVariableTable":
		return &ATTRIBUTE_LOCAL_VARIABLE_TABLE{}
	case "SourceFile":
		return &ATTRIBUTE_SOURCE_FILE {
			constantPool: constantPool,
		}
	case "Synthetic":
		return &ATTRIBUTE_SYNTHETIC{}
	default:
		return &UnparsedAttribute{
			name: attrName,
			length: attrLength,
			info: nil,
		}
	}
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据读入两个字节作为属性数, 创建 []AttributeInfo, 对其进行遍历赋值
func resolveAttributes(reader *ClassReader, constantPool ConstantPool) []AttributeInfo {
	attributeCount := reader.readU2()
	attributes := make([]AttributeInfo, attributeCount)
	for index := range attributes {
		attributes[index] = readAttribute(reader, constantPool)
	}
	return attributes
}

// 读入属性名, 属性长度
func readAttribute(reader *ClassReader, constantPool ConstantPool) AttributeInfo {
	attrNameIndex := reader.readU2()
	attrName := constantPool.getUtf8(attrNameIndex)
	attrLength := reader.readU4()
	attrInfo := newAttributeInfo(attrName, attrLength, constantPool)
	attrInfo.readInfo(reader)
	return attrInfo
}