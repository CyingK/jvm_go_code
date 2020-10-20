package classfile

// 属性（注解）
type ATTRIBUTE_MARKER struct {
}

//--------------------------------------------------------------------功能类方法

// 啥都不干
func (self *ATTRIBUTE_MARKER) readInfo(reader *ClassReader) {
}

// @Deprecated
type ATTRIBUTE_DEPRECATED struct {
	ATTRIBUTE_MARKER
}

// @Synthetic
type ATTRIBUTE_SYNTHETIC struct {
	ATTRIBUTE_MARKER
}
