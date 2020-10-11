package classfile

import "log"

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
	log.Println("\t    (4Byte)CONSTANT_integer_info")
	log.Println("\t\tvalue：", self.val)
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
	log.Println("\t    (4Byte)CONSTANT_float_info")
	log.Println("\t\tvalue：", self.val)
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
	log.Println("\t    (8Byte)CONSTANT_long_info")
	log.Println("\t\tvalue：", self.val)
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
	log.Println("\t    (8Byte)CONSTANT_double_info")
	log.Println("\t\tvalue：", self.val)
}