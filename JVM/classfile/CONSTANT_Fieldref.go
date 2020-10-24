package classfile

/*
	CONSTANT_Fieldref_info {
		u1 tag;
		u2 class_index;
		u2 name_and_type_index;
	}
*/

// 字段类型常量
type CONSTANT_Fieldref_info struct {
	CONSTANT_Memberref_info
}
