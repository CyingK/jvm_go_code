package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

/*
 * 接口和抽象类不能实例化
 */
func (self *NEW) Execute(frame *rtda.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	classRef := constantPool.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.GetOperandStack().PushRef(ref)
}

func (self *NEW) String() string {
	return "{type：new; " + self.Index16Instruction.String() + "}\t"
}