package heap

import "jvm_go_code/JVM/classfile"

type MemberRef struct {
	SymRef
	name		string
	descriptor	string
}

func (self *MemberRef) copyMemberRefInfo(refInfo *classfile.CONSTANT_Memberref_info) {
	self.className = refInfo.GetClassName()
	self.name, self.descriptor = refInfo.GetNameAndDescriptor()
}

func (self *MemberRef) GetName() string {
	return self.name
}

func (self *MemberRef) Descriptor() string {
	return self.descriptor
}