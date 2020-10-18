package heap

import "jvm_go_code/invoke_return/classfile"

type InterfaceMethodRef struct {
	MemberRef
	method		*Method
}

func newInterfaceMethodRef(constantpool *ConstantPool, refInfo *classfile.CONSTANT_INTERFACE_METHOD_REF_INFO) *InterfaceMethodRef {
	ref := &InterfaceMethodRef{}
	ref.constantPool = constantpool
	ref.copyMemberRefInfo(&refInfo.CONSTANT_MEMBER_REF_INFO)
	return ref
}