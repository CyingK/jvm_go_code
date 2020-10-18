package references

import (
	"fmt"
	"jvm_go_code/invoke_return/instructions/base"
	"jvm_go_code/invoke_return/rtda"
	"jvm_go_code/invoke_return/rtda/heap"
)

type NEW struct {
	base.Index16Instruction
}

/*
 * 接口和抽象类不能实例化
 */
func (self *NEW) Execute(frame *rtda.Frame) {
	constantPool := frame.Method().Class().ConstantPool()
	classRef := constantPool.GetConstant(self.Index).(*heap.ClassRef)
	class := classRef.ResolvedClass()
	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}
	ref := class.NewObject()
	frame.OperandStack().PushRef(ref)
	fmt.Printf("new: 从常量池获取到%v的引用信息, 然后装载这个类并创建了一个该类对象推入操作数栈\n", classRef.GetClassName())
}

func (self *NEW) String() string {
	return "{type：new; " + self.Index16Instruction.String() + "}"
}