package classfile

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

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	c := newConstantInfo(tag, cp)
	c.readInfo(reader)
	return c
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
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
		return &CONSTANT_STRING_INFO{}
	case CONSTANT_Class:
		return &CONSTANT_CLASS_INFO{}
	case CONSTANT_Fieldref:
		return &CONSTANT_FIELD_REF_INFO {
		CONSTANT_MEMBER_REF_INFO {
				cp: cp,
			},
		}
	case CONSTANT_Methodref:
		return &CONSTANT_METHOD_REF_INFO {
		CONSTANT_MEMBER_REF_INFO {
				cp: cp,
			},
		}
	case CONSTANT_InterfaceMethodref:
		return &CONSTANT_INTERFACE_METHOD_REF_INFO {
		CONSTANT_MEMBER_REF_INFO {
				cp: cp,
			},
		}
	case CONSTANT_NameAndType:
		return &CONSTANT_NAME_AND_TYPE_INFO{}
	default:
		panic("java.lang.ClassFormatError: Constant Pool Tag !")
	}
}