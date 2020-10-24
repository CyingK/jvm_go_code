package classfile

/*
	CONSTANT_InterfaceMethodref_info {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}
*/

// 接口方法类型常量
type CONSTANT_InterfaceMethodref_info struct {
	CONSTANT_Memberref_info
}
