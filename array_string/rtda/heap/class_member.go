package heap

import "jvm_go_code/array_string/classfile"

type ClassMember struct {
	accessFlags		uint16		// 访问标识
	name			string		// 方法/属性名称
	descriptor		string		// 方法/属性描述
	class			*Class		// 归属类
}

// 复制成员信息, 包括 accessFlags, name, descriptor
func (self *ClassMember) copyMemberInfo(memberInfo *classfile.MemberInfo) {
	self.accessFlags = memberInfo.GetAccessFlags()
	self.name = memberInfo.GetName()
	self.descriptor = memberInfo.GetDescriptor()
}

// 如果该类的权限是 public, 返回 true, 否则返回 false
func (self *ClassMember) IsPublic() bool {
	return 0 != self.accessFlags & ACC_PUBLIC
}

// 如果该类的权限是 private, 返回 true, 否则返回 false
func (self *ClassMember) IsPrivate() bool {
	return 0 != self.accessFlags * ACC_PRIVATE
}

// 如果该类的权限是 protected, 返回 true, 否则返回 false
func (self *ClassMember) IsProtected() bool {
	return 0 != self.accessFlags & ACC_PROTECTED
}

// 如果该类的权限是 static, 返回 true, 否则返回 false
func (self *ClassMember) IsStatic() bool {
	return 0 != self.accessFlags & ACC_STATIC
}

// 如果该类的权限是 final, 返回 true, 否则返回 false
func (self *ClassMember) IsFinal() bool {
	return 0 != self.accessFlags & ACC_FINAL
}

// 如果该类的权限是 synthetic, 返回 true, 否则返回 false
func (self *ClassMember) IsSynthetic() bool {
	return 0 != self.accessFlags & ACC_SYNTHETIC
}

// 获取 name
func (self *ClassMember) GetName() string {
	return self.name
}

// 获取 descriptor
func (self *ClassMember) Descriptor() string {
	return self.descriptor
}

// 获取 class
func (self *ClassMember) GetClass() *Class {
	return self.class
}

// 判断 otherClass 是否能访问 self, 首先如果 self 的权限为 public 直接返回 true. 其次如果 self.class是 protected, 要么 otherClass 和 self.class 相等, 要么 otherClass 是 self.class 的子类,
// 要么 otherClass 和 self.class 的包路径相同, 以上三种情况返回 true. 最后如果 self.class 的权限是默认, 比较 otherClass 和 self.class 的包路径, 相同则返回 true, 否则返回 otherClass == thisClass
func (self *ClassMember) isAccessibleTo(otherClass *Class) bool {
	if self.IsPublic() {
		return true
	}
	thisClass := self.class
	if self.IsProtected() {
		return otherClass == thisClass || otherClass.IsSubClassOf(thisClass) || thisClass.GetPackageName() == otherClass.GetPackageName()
	}
	if !self.IsPrivate() {
		return thisClass.GetPackageName() == otherClass.GetPackageName()
	}
	return otherClass == thisClass
}