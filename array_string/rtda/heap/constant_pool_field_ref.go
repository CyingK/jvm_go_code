package heap

import "jvm_go_code/array_string/classfile"

type FieldRef struct {
	MemberRef
	field *Field
}

// 创建新的 FieldRef, 初始化 constantPool, 并调用 copyMemberRefInfo 进行属性信息的复制
func newFieldRef(constantPool *ConstantPool, refInfo *classfile.CONSTANT_FIELD_REF_INFO) *FieldRef {
	ref := &FieldRef{}
	ref.constantPool = constantPool
	ref.copyMemberRefInfo(&refInfo.CONSTANT_MEMBER_REF_INFO)
	return ref
}

// 获取 field
func (self *FieldRef) ResolvedField() *Field {
	if self.field == nil {
		self.resolveFieldRef()
	}
	return self.field
}

// 在归属类中查找是否有 self.name, self.descriptor 一致的属性, 并判断该属性能否访问运行时常量池的归属类
func (self *FieldRef) resolveFieldRef() {
	constantPoolClass := self.constantPool.class
	selfClass := self.ResolvedClass()
	field := lookupField(selfClass, self.name, self.descriptor)
	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}
	if !field.isAccessibleTo(constantPoolClass) {
		panic("java.lang.IllegalAccessError")
	}
	self.field = field
}

// 分别在 class, class 的所有接口, class 的所有父类中查找与 name, descriptor 一致的属性并返回
func lookupField(class *Class, name string, descriptor string) *Field {
	for _, item := range class.fields {
		if item.name == name && item.descriptor == descriptor {
			return item
		}
	}
	for _, item := range class.interfaces {
		if field := lookupField(item, name, descriptor); field != nil {
			return field
		}
	}
	if class.superClass != nil {
		return lookupField(class.superClass, name, descriptor)
	}
	return nil
}