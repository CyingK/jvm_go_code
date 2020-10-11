package classfile

import (
	"fmt"
	"log"
)

/*
 * 负责模块化读入 class 文件，检查 magic, version, 对外提供 get 方法
 */

type ClassFile struct {
	magic			uint32				// 魔数
	minorVersion	uint16				// 次版本号
	majorVersion	uint16				// 主版本号
	constantPool	ConstantPool		// 常量池
	accessFlags		uint16				// 访问标志
	thisClass		uint16				// 本类标识
	superClass		uint16				// 超类标识
	interfaces		[]uint16			// 接口标识
	fields			[]*MemberInfo		// 字段标识
	methods			[]*MemberInfo		// 方法标识
	atrributes		[]AttributeInfo	// 属性标识
}


func Parse(classData []byte) (classFile *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			if err, ok = r.(error); !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	classReader := &ClassReader{
		data: classData,
	}
	classFile = &ClassFile{}
	classFile.read(classReader)
	return
}

/*
 * 读入 class 文件
 */
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.atrributes = readAttributes(reader, self.constantPool)
}

/*
 * 检验魔数 0xCAFEBABE
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	if self.magic = reader.readUint32(); self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/*
 * 检验版本是否在可运行范围内
 */
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedVersionError")
}

/*
 * 获取次版本
 */
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

/*
 * 获取主版本
 */
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

/*
 * 获取常量池
 */
func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

/*
 * 获取访问标识
 */
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

/*
 * 获取所有字段名
 */
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

/*
 * 获取所有方法名
 */
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}

/*
 * 获取类名
 */
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

/*
 * 获取超类名
 */
func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

/*
 * 获取所有接口名
 */
func (self *ClassFile) InterfaceNames() []string {
	interfacesNames := make([]string, len(self.interfaces))
	for index, item := range self.interfaces {
		interfacesNames[index] = self.constantPool.getClassName(item)
	}
	return interfacesNames
}