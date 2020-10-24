package classfile

/*
	CONSTANT_Double_info {
		u1 tag;
		u4 high_bytes;
		u4 low_bytes;
	}
 */

// 双精度浮点型常量
type CONSTANT_Double_info struct {
	val float64	// 值
}

// 从二进制数据中读取 8 个字节，存入 value
func (self *CONSTANT_Double_info) readInfo(reader *ClassReader) {
	bytes := reader.readU8()
	self.val = float64(bytes)
}

// 获取 value
func (self *CONSTANT_Double_info) GetValue() float64 {
	return self.val
}
