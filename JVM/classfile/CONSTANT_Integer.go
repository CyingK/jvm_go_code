package classfile

/*
	CONSTANT_Integer_info {
		u1 tag;
		u4 bytes;
	}
 */

// 整形常量
type CONSTANT_Integer_info struct {
	value int32 // 值
}

// 从二进制数据中读取 4 个字节，存入 value
func (self *CONSTANT_Integer_info) readInfo(reader *ClassReader) {
	bytes := reader.readU4()
	self.value = int32(bytes)
}

// 获取 value
func (self *CONSTANT_Integer_info) GetValue() int32 {
	return self.value
}