package classfile

import "encoding/binary"

/*
 * 只是 []byte 的一个包装, 提供便于读取字节的方法
 */

type ClassReader struct {
	data	[]byte
}

/*
 * 从 class 文件读入 1 个字节
 */
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

/*
 * 从 class 文件读入 2 个字节
 */
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

/*
 * 从 class 文件读入 4 个字节
 */
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

/*
 * 从 class 文件读入 8 个字节
 */
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/*
 * 从 class 文件读入 count 次 2 个字节
 */
func (self *ClassReader) readUint16s() []uint16 {
	count := self.readUint16()
	results := make([]uint16, count)
	for index, _ := range results {
		results[index] = self.readUint16()
	}
	return results
}

/*
 * 从 class 文件读入 length 个字节
 */
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}