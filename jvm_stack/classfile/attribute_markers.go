package classfile

type ATTRIBUTE_MARKER struct {
	
}

func (self *ATTRIBUTE_MARKER) readInfo(reader *ClassReader) {
	// 啥也不做
}

type ATTRIBUTE_DEPRECATED struct {
	ATTRIBUTE_MARKER
}

type ATTRIBUTE_SYNTHETIC struct {
	ATTRIBUTE_MARKER
}
