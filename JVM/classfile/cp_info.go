package classfile

/*
 * 根据 tag 生成对应的 ConstantInfo 并返回
 */

const (
	CONSTANT_Utf8					= 1		// 具体的字符串
	CONSTANT_Integer				= 3		// 整型
	CONSTANT_Float					= 4		// 浮点型
	CONSTANT_Long					= 5		// 长整型
	CONSTANT_Double					= 6		// 双精度浮点型
	CONSTANT_Class					= 7		// 类
	CONSTANT_String					= 8		// 字符串引用
	CONSTANT_Fieldref				= 9		// 字段引用
	CONSTANT_Methodref				= 10	// 方法引用
	CONSTANT_InterfaceMethodref		= 11	// 接口方法引用
	CONSTANT_NameAndType			= 12	// 名称和描述
	CONSTANT_MethodHandle       	= 15
	CONSTANT_MethodType         	= 16
	CONSTANT_InvokeDynamic      	= 18
)

// 常量信息
type ConstantInfo interface {
	readInfo(reader *ClassReader)
}

//--------------------------------------------------------------------构造器

// 根据 tag，生成不同的 ConstantInfo 返回
func newConstantInfo(tag U1, constantPool ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &CONSTANT_Integer_info{}
	case CONSTANT_Float:
		return &CONSTANT_Float_info{}
	case CONSTANT_Long:
		return &CONSTANT_Long_info{}
	case CONSTANT_Double:
		return &CONSTANT_Double_info{}
	case CONSTANT_Utf8:
		return &CONSTANT_Utf8_info{}
	case CONSTANT_String:
		return &CONSTANT_String_info{
			constantPool: constantPool,
		}
	case CONSTANT_Class:
		return &CONSTANT_Class_info{
			tag: 			tag,
			constantPool: 	constantPool,
		}
	case CONSTANT_Fieldref:
		return &CONSTANT_Fieldref_info{
			CONSTANT_Memberref_info{
				constantPool: constantPool,
			},
		}
	case CONSTANT_Methodref:
		return &CONSTANT_Methodref_info{
			CONSTANT_Memberref_info{
				constantPool: constantPool,
			},
		}
	case CONSTANT_InterfaceMethodref:
		return &CONSTANT_InterfaceMethodref_info{
			CONSTANT_Memberref_info{
				constantPool: constantPool,
			},
		}
	case CONSTANT_NameAndType:
		return &CONSTANT_NameAndType_info{}
	case CONSTANT_MethodType:
		return &CONSTANT_MethodType_info{}
	case CONSTANT_MethodHandle:
		return &CONSTANT_MethodHandle_info{}
	case CONSTANT_InvokeDynamic:
		return &CONSTANT_InvokeDynamic_info{}
	default:
		panic("java.lang.ClassFormatError: Constant Pool Tag !")
	}
}

//--------------------------------------------------------------------功能类方法

// 读入 tag 并生成一个 ConstantInfo 返回
func readConstantInfo(reader *ClassReader, constantPool ConstantPool) ConstantInfo {
	tag := reader.readU1()
	constantInfo := newConstantInfo(tag, constantPool)
	constantInfo.readInfo(reader)
	return constantInfo
}
