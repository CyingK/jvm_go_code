package classfile

import "log"

/*
 * utf8 类型 tag
 */
type CONSTANT_UTF8_INFO struct {
	str string	// 字符串
}

/*
 * 从数据中读取 2 个字节，作为 length，然后从 数据中读取 length 个字节，存入 str
 */
func (self *CONSTANT_UTF8_INFO) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
	log.Println("\t    (4Byte)CONSTANT_utf_info")
	log.Println("\t\tstr：", self.str)
}

/*
 * 反编码 UTF8
 */
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
