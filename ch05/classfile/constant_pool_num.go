package classfile

/*
 * integer 类型 tag
 */
type CONSTANT_INTEGER_INFO struct {
	val int32	// 值
}

/*
 * 从数据中读取 4 个字节，存入 val
 */
func (self *CONSTANT_INTEGER_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

/*
 * float 类型 tag
 */
type CONSTANT_FLOAT_INFO struct {
	val float32	// 值
}

/*
 * 从数据中读取 4 个字节，存入 val
 */
func (self *CONSTANT_FLOAT_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = float32(bytes)
}

/*
 * long 类型 tag
 */
type CONSTANT_LONG_INFO struct {
	val int64	// 值
}

/*
 * 从数据中读取 8 个字节，存入 val
 */
func (self *CONSTANT_LONG_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

/*
 * double 类型 tag
 */
type CONSTANT_DOUBLE_INFO struct {
	val float64	// 值
}

/*
 * 从数据中读取 8 个字节，存入 val
 */
func (self *CONSTANT_DOUBLE_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = float64(bytes)
}