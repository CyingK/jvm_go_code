package classfile

type CONSTANT_INTEGER_INFO struct {
	val int32
}

func (self *CONSTANT_INTEGER_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = int32(bytes)
}

type CONSTANT_FLOAT_INFO struct {
	val float32
}

func (self *CONSTANT_FLOAT_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint32()
	self.val = float32(bytes)
}

type CONSTANT_LONG_INFO struct {
	val int64
}

func (self *CONSTANT_LONG_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = int64(bytes)
}

type CONSTANT_DOUBLE_INFO struct {
	val float64
}

func (self *CONSTANT_DOUBLE_INFO) readInfo(reader *ClassReader) {
	bytes := reader.readUint64()
	self.val = float64(bytes)
}