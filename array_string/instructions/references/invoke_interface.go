package references

import (
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
	"strconv"
)

// 当通过 invokeinterface 指令调用接口方法时, 因为 this 引用可以指向任何实现了该接口的类的实例，所以无法使用 vtable 技术
type INVOKE_INTERFACE struct {
	index uint
}

func (self *INVOKE_INTERFACE) GetOperands(reader *base.ByteCodeReader) {
	self.index = uint(reader.ReadInt16())
	reader.ReadUint8()
	reader.ReadUint8()
}

func (self *INVOKE_INTERFACE) Execute(frame *rtda.Frame) {
	// 从运行时常量池中拿到并解析接口方法符号引用，如果解析后的方法是静态方法或私有方法，则抛出异常
	constangPool := frame.GetMethod().GetClass().GetConstantPool()
	methodRef := constangPool.GetConstant(self.index).(*heap.InterfaceMethodRef)
	resolvedMethod := methodRef.ResolvedInterfaceMethod()
	if resolvedMethod.IsStatic() || resolvedMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 从操作数栈中弹出this引用，如果引用是null，则抛出异常
	ref := frame.GetOperandStack().GetRefFromTop(resolvedMethod.GetArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	// 如果引用所指对象的类没有实现解析出来的接口，则抛出异常
	if !ref.GetClass().IsImplements(methodRef.ResolvedClass()) {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 查找最终要调用的方法。如果找不到，或者找到的方法是抽象的，则抛出异常。
	methodToBeInvoked := heap.LookupMethodInClass(ref.GetClass(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	// 如果找到的方法不是public, 则抛出异常
	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.IllegalAccessError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func (self *INVOKE_INTERFACE) String() string {
	return "{type：invoke_interface; index: " + strconv.Itoa(int(self.index)) + "}"
}