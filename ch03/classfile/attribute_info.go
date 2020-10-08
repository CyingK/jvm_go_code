package classfile

import "log"

type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func readAttributes(reader *ClassReader, constantPool ConstantPool) []AttributeInfo {
	attributeCount := reader.readUint16()
	log.Println("\t\t(2Byte)读入属性个数：", attributeCount)
	attributes := make([]AttributeInfo, attributeCount)
	for index := range attributes {
		log.Printf("\t\t    [%d]", index + 1)
		attributes[index] = readAttribute(reader, constantPool)
	}
	return attributes
}

func readAttribute(reader *ClassReader, constantPool ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUint16()
	attrName := constantPool.getUtf8(attrNameIndex)
	log.Println("\t\t\t(2Byte)名称下标：", attrNameIndex, "[", attrName, "]")
	attrLength := reader.readUint32()
	log.Println("\t\t\t(4Byte)长度：", attrLength)
	attrInfo := newAttributeInfo(attrName, attrLength, constantPool)
	attrInfo.readInfo(reader)
	return attrInfo
}

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

