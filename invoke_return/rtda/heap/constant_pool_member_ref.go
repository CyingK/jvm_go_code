package heap

import "jvm_go_code/invoke_return/classfile"

type MemberRef struct {
	SymRef
	name		string
	descriptor	string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.CONSTANT_MEMBER_REF_INFO) {
	self.className = refInfo.ClassName()
	self.name, self.descriptor = refInfo.NameAndDescriptor()
}

func (self *MemberRef) Name() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}