package classfile

/*
 * 根据 tag 生成对应的 ConstantInfo 并返回
 */

const (
	CONSTANT_Utf8					= 1
	CONSTANT_Integer				= 3
	CONSTANT_Float				= 4
	CONSTANT_Long					= 5
	CONSTANT_Double				= 6
	CONSTANT_Class				= 7
	CONSTANT_String				= 8
	CONSTANT_Fieldref				= 9
	CONSTANT_Methodref			= 10
	CONSTANT_InterfaceMethodref	= 11
	CONSTANT_NameAndType			= 12
)

type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

/*
 * 读入一个 ConstantInfo
 */
func readConstantInfo(reader *ClassReader, constantPool ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	constantInfo := newConstantInfo(tag, constantPool)
	constantInfo.readInfo(reader)
	return constantInfo
}

/*
 * 根据 tag，生成不同的 ConstantInfo 返回
 */
func newConstantInfo(tag uint8, constantPool ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &CONSTANT_INTEGER_INFO{}
	case CONSTANT_Float:
		return &CONSTANT_FLOAT_INFO{}
	case CONSTANT_Long:
		return &CONSTANT_LONG_INFO{}
	case CONSTANT_Double:
		return &CONSTANT_DOUBLE_INFO{}
	case CONSTANT_Utf8:
		return &CONSTANT_UTF8_INFO{}
	case CONSTANT_String:
		return &CONSTANT_STRING_INFO{
		constantPool: constantPool,
		}
	case CONSTANT_Class:
		return &CONSTANT_CLASS_INFO{
		constantPool: constantPool,
		}
	case CONSTANT_Fieldref:
		return &CONSTANT_FIELD_REF_INFO {
		CONSTANT_MEMBER_REF_INFO {
				constantPool: constantPool,
			},
		}
	case CONSTANT_Methodref:
		return &CONSTANT_METHOD_REF_INFO {
		CONSTANT_MEMBER_REF_INFO {
				constantPool: constantPool,
			},
		}
	case CONSTANT_InterfaceMethodref:
		return &CONSTANT_INTERFACE_METHOD_REF_INFO {
		CONSTANT_MEMBER_REF_INFO {
				constantPool: constantPool,
			},
		}
	case CONSTANT_NameAndType:
		return &CONSTANT_NAME_AND_TYPE_INFO{}
	default:
		panic("java.lang.ClassFormatError: Constant Pool Tag !")
	}
}