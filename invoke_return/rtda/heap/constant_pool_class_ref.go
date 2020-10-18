package heap

import "jvm_go_code/invoke_return/classfile"

type ClassRef struct {
	SymRef
}

// 创建新的 ClassRef, 初始化 constantPool, className 并返回
func newClassRef(constantPool *ConstantPool, classInfo *classfile.CONSTANT_CLASS_INFO) *ClassRef {
	ref := &ClassRef{}
	ref.constantPool = constantPool
	ref.className = classInfo.Name()
	return ref
}