package classfile

import "log"

/*
 * class 类型 tag
 */
type CONSTANT_CLASS_INFO struct {
	constantPool ConstantPool	// 常量池
	nameIndex    uint16			// 名称索引
}

/*
 * 从数据中读取 2 个字节，作为名称索引
 */
func (self *CONSTANT_CLASS_INFO) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	log.Printf("\t    (2Byte)CONSTANT_class_info")
	log.Println("\t\tindex：", self.nameIndex)
}

/*
 * 根据 name 索引，在常量池中找到并返回 name
 */
func (self *CONSTANT_CLASS_INFO) Name() string {
	return self.constantPool.getUtf8(self.nameIndex)
}