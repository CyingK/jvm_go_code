package classfile

import (
	"fmt"
)

/*
 * 负责模块化读入 class 文件，检查 magic, version, 对外提供 get 方法, 简单来说就是把二进制数据结构化成一个 ClassFile 结构体
 */

// 类文件结构
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

//--------------------------------------------------------------------Getters

// 获取次版本
func (self *ClassFile) GetMinorVersion() uint16 {
	return self.minorVersion
}

// 获取主版本
func (self *ClassFile) GetMajorVersion() uint16 {
	return self.majorVersion
}

// 获取常量池
func (self *ClassFile) GetConstantPool() ConstantPool {
	return self.constantPool
}

// 获取访问标识
func (self *ClassFile) GetAccessFlags() uint16 {
	return self.accessFlags
}

// 获取所有字段名
func (self *ClassFile) GetFields() []*MemberInfo {
	return self.fields
}

// 获取所有方法名
func (self *ClassFile) GetMethods() []*MemberInfo {
	return self.methods
}

// 获取类名
func (self *ClassFile) GetClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

// 获取超类名
func (self *ClassFile) GetSuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return ""
}

// 获取所有接口名
func (self *ClassFile) GetInterfaceNames() []string {
	interfacesNames := make([]string, len(self.interfaces))
	for index, item := range self.interfaces {
		interfacesNames[index] = self.constantPool.getClassName(item)
	}
	return interfacesNames
}

//--------------------------------------------------------------------功能类方法

// 创建 class 整体结构
func ResolveClassData(classData []byte) (classFile *ClassFile, err error) {
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
	classFile.resolve(classReader)
	return
}

// 按照 class 的结构规范解析二进制数据
func (self *ClassFile) resolve(reader *ClassReader) {
	// part 1 magic
	self.resolveAndCheckMagic(reader)
	// part 2 minor_version
	// part 3 major_version
	self.resolveAndCheckVersion(reader)
	// part 4 constant_pool_count
	// part 5 constant_pool[constant_pool_count - 1]
	self.constantPool = resolveConstantPool(reader)
	// part 6 access_flags
	self.accessFlags = reader.readUint16()
	// part 7 this_class
	self.thisClass = reader.readUint16()
	// part 8 super_class
	self.superClass = reader.readUint16()
	// part 9 interfaces_count
	// part 10 interfaces[interfaces_count]
	self.interfaces = reader.readUint16s()
	// part 11 fileds_count
	// part 12 fileds[fields_count]
	self.fields = resolveMembers(reader, self.constantPool)
	// part 13 methods_count
	// part 14 methods[methods_count]
	self.methods = resolveMembers(reader, self.constantPool)
	// part 15 attributes_count
	// part 16 attributes[attributes_count]
	self.atrributes = resolveAttributes(reader, self.constantPool)
}

// 解析魔数
func (self *ClassFile) resolveAndCheckMagic(reader *ClassReader) {
	self.magic = reader.readUint32()
	if self.magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

// 解析版本
func (self *ClassFile) resolveAndCheckVersion(reader *ClassReader) {
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