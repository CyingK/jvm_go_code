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

func (self *CONSTANT_INTEGER_INFO) Value() int32 {
	return self.val
}

/*
 * float 类型 tag
 */
type CONSTANT_FLOAT_INFO struct {
	val float32	// 值
}

func (self *CONSTANT_FLOAT_INFO) Value() float32 {
	return self.val
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

func (self *CONSTANT_LONG_INFO) Value() int64 {
	return self.val
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

func (self *CONSTANT_DOUBLE_INFO) Value() float64 {
	return self.val
}