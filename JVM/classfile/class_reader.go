package classfile

import "encoding/binary"

/*
 * 只是 []byte 的一个包装, 提供便于读取字节的方法
 */

// 二进制数据读取工具
type ClassReader struct {
	data	[]byte		// 二进制数据
}

//--------------------------------------------------------------------Getters

// 读入异常表长度, 据此创建异常表, 然后遍历每一个元素, 对其进行二进制数据的读入填充, 最后返回此异常表
func getExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readU2()
	exceptionTables := make([]*ExceptionTableEntry, exceptionTableLength)
	for index := range exceptionTables {
		exceptionTables[index] = &ExceptionTableEntry {
			startPc: reader.readU2(),
			endPc: reader.readU2(),
			handlerPc: reader.readU2(),
			catchType: reader.readU2(),
		}
	}
	return exceptionTables
}


//--------------------------------------------------------------------功能类方法

// 从 class 文件读入 1 个字节
func (self *ClassReader) readU1() U1 {
	val := self.data[0]
	self.data = self.data[1:]
	return U1(val)
}

// 从 class 文件读入 2 个字节
func (self *ClassReader) readU2() U2 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return U2(val)
}

// 从 class 文件读入 4 个字节
func (self *ClassReader) readU4() U4 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return U4(val)
}

// 从 class 文件读入 8 个字节
func (self *ClassReader) readU8() U8 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return U8(val)
}

// 从 class 文件读入 count 次 2 个字节
func (self *ClassReader) readMultiU2() []U2 {
	count := self.readU2()
	results := make([]U2, count)
	for index, _ := range results {
		results[index] = self.readU2()
	}
	return results
}

// 从 class 文件读入 length 个字节
func (self *ClassReader) readUn(n U4) []byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}