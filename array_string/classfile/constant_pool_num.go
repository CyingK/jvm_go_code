package classfile

// 整形常量
type CONSTANT_INTEGER_INFO struct {
	value int32 // 值
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据中读取 4 个字节，存入 value
func (self *CONSTANT_INTEGER_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.value = int32(bytes)
}

//--------------------------------------------------------------------Getters

// 获取 value
func (self *CONSTANT_INTEGER_INFO) GetValue() int32 {
	return self.value
}

// 浮点型常量
type CONSTANT_FLOAT_INFO struct {
	val float32	// 值
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据中读取 4 个字节，存入 value
func (self *CONSTANT_FLOAT_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = float32(bytes)
}

//--------------------------------------------------------------------Getters

// 获取 value
func (self *CONSTANT_FLOAT_INFO) GetValue() float32 {
	return self.val
}

// 长整型常量
type CONSTANT_LONG_INFO struct {
	val int64	// 值
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据中读取 8 个字节，存入 value
func (self *CONSTANT_LONG_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

//--------------------------------------------------------------------Getters

// 获取 value
func (self *CONSTANT_LONG_INFO) GetValue() int64 {
	return self.val
}

// 双精度浮点型常量
type CONSTANT_DOUBLE_INFO struct {
	val float64	// 值
}

//--------------------------------------------------------------------功能类方法

// 从二进制数据中读取 8 个字节，存入 value
func (self *CONSTANT_DOUBLE_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = float64(bytes)
}

//--------------------------------------------------------------------Getters

// 获取 value
func (self *CONSTANT_DOUBLE_INFO) GetValue() float64 {
	return self.val
}