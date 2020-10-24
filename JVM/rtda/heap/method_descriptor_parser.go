package heap

import (
	"strings"
)

type MethodDescriptorParser struct {
	raw		string					// 原始数据
	offset	int						// 偏移量
	parsed	*MethodDescriptor		// 解析完成的
}

// 解析方法的描述
func parseMethodDescriptor(descriptor string) *MethodDescriptor {
	parser := &MethodDescriptorParser{}
	return parser.parse(descriptor)
}

// 分别调用方法进行各阶段解析
func (self *MethodDescriptorParser) parse(descriptor string) *MethodDescriptor {
	self.raw = descriptor
	self.parsed = &MethodDescriptor{}
	self.startParams()
	self.parseParamTypes()
	self.endParams()
	self.parseReturnType()
	self.finish()
	return self.parsed
}

// 判断首字符是不是(
func (self *MethodDescriptorParser) startParams() {
	flag := self.readUint8()
	if flag != '(' {
		self.causePanic()
	}
}

// 判断尾字符是不是)
func (self *MethodDescriptorParser) endParams() {
	flag := self.readUint8()
	if flag != ')' {
		self.causePanic()
	}
}

// 解析每个参数的类型描述
func (self *MethodDescriptorParser) parseParamTypes() {
	for {
		_type_ := self.parseFieldType()
		if _type_ != "" {
			self.parsed.addParameterType(_type_)
		} else {
			break
		}
	}
}

// 解析返回值类型
func (self *MethodDescriptorParser) parseReturnType() {
	flag := self.readUint8()
	if flag == 'V' {
		self.parsed.returnType = "V"
		return
	}
	self.unreadUint8()
	_type_ := self.parseFieldType()
	if _type_ != "" {
		self.parsed.returnType = _type_
		return
	}
	self.causePanic()
}

// 判断 self.offset 读完后是否与 len(self.raw) 相同
func (self *MethodDescriptorParser) finish() {
	if self.offset != len(self.raw) {
		self.causePanic()
	}
}

// 异常中断
func (self *MethodDescriptorParser) causePanic() {
	panic("BAD descriptor: " + self.raw)
}

// 截取原始数据中 offset 处的字符并返回
func (self *MethodDescriptorParser) readUint8() uint8 {
	symbol := self.raw[self.offset]
	self.offset++
	return symbol
}

// 偏移量往后退一个元素
func (self *MethodDescriptorParser) unreadUint8() {
	self.offset--
}

// 根据读入的字符, 返回对应的类型字符串
func (self *MethodDescriptorParser) parseFieldType() string {
	flag := self.readUint8()
	switch flag {
	case 'B':
		return "B"
	case 'C':
		return "C"
	case 'D':
		return "D"
	case 'F':
		return "F"
	case 'I':
		return "I"
	case 'J':
		return "J"
	case 'S':
		return "S"
	case 'Z':
		return "Z"
	case 'L':
		return self.parseObjectType()
	case '[':
		return self.parseArrayType()
	default:
		self.unreadUint8()
		return ""
	}
}

// 截取对象的类型描述
func (self *MethodDescriptorParser) parseObjectType() string {
	unread := self.raw[self.offset:]
	semicoloIndex := strings.IndexRune(unread, ';')
	if semicoloIndex == -1 {
		self.causePanic()
		return ""
	} else {
		objStart := self.offset - 1
		objEnd := self.offset + semicoloIndex + 1
		self.offset = objEnd
		descriptor := self.raw[objStart:objEnd]
		return descriptor
	}
}

// 截取数组的类型描述
func (self *MethodDescriptorParser) parseArrayType() string {
	arrayStart := self.offset - 1
	self.parseFieldType()
	arrayEnd := self.offset
	descriptor := self.raw[arrayStart:arrayEnd]
	return descriptor
}