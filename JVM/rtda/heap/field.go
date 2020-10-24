package heap

import "jvm_go_code/JVM/classfile"

type Field struct {
	ClassMember
	constValueIndex	uint
	slotId			uint
}

func newFields(class *Class, classFileFields []*classfile.MemberInfo) []*Field {
	fields := make([]*Field, len(classFileFields))
	for index, item := range classFileFields {
		fields[index] = &Field{}
		fields[index].class = class
		fields[index].copyMemberInfo(item)
		fields[index].copyAttributes(item)
	}
	return fields
}

func (self *Field) copyAttributes(classFileField *classfile.MemberInfo) {
	if varAttr := classFileField.GetConstantValueAttribute(); varAttr != nil {
		self.constValueIndex = uint(varAttr.GetConstantValueIndex())
	}
}

func (self *Field) IsVolatile() bool {
	return 0 != self.accessFlags & ACC_VOLATILE
}

func (self *Field) IsTransient() bool {
	return 0 != self.accessFlags & ACC_TRANSIENT
}

func (self *Field) IsEnum() bool {
	return 0 != self.accessFlags & ACC_ENUM
}

func (self *Field) ConstValueIndex() uint {
	return self.constValueIndex
}

func (self *Field) SlotId() uint {
	return self.slotId
}

func (self *Field) isLongOrDouble() bool {
	return self.descriptor == "J" || self.descriptor == "D"
}

func (self *Field) GetType() *Class {
	className := toClassName(self.descriptor)
	return self.class.loader.LoadClass(className)
}

func (self *Field) GetAccessFlags() uint16 {
	return self.accessFlags
}