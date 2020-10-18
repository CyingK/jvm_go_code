package heap

import "jvm_go_code/invoke_return/classfile"

type MethodRef struct {
	MemberRef
	method 		*Method
}

func newMethodRef(constantPool *ConstantPool, refInfo *classfile.CONSTANT_METHOD_REF_INFO) *MethodRef {
	ref := &MethodRef{}
	ref.constantPool = constantPool
	ref.copyMemberRefInfo(&refInfo.CONSTANT_MEMBER_REF_INFO)
	return ref
}