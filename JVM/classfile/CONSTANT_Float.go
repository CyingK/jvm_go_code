package classfile

/*
	CONSTANT_Float_info {
		 u1 tag;
		 u4 bytes;
	}
 */

// 浮点型常量
type CONSTANT_Float_info struct {
	val float32	// 值
}

// 从二进制数据中读取 4 个字节，存入 value
func (self *CONSTANT_Float_info) readInfo(reader *ClassReader) {
	bytes := reader.readU4()
	self.val = float32(bytes)
}

// 获取 value
func (self *CONSTANT_Float_info) GetValue() float32 {
	return self.val
}
