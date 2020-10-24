package classfile

/*
	CONSTANT_Long_info {
		u1 tag;
		u4 high_bytes;
		u4 low_bytes;
	}
 */

// 长整型常量
type CONSTANT_Long_info struct {
	val 	U8	// 值
}

// 从二进制数据中读取 8 个字节，存入 value
func (self *CONSTANT_Long_info) readInfo(reader *ClassReader) {
	bytes := reader.readU8()
	self.val = bytes
}

// 获取 value
func (self *CONSTANT_Long_info) GetValue() U8 {
	return self.val
}
