package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

// 使用 super 关键字调用超类中的方法不能使用invokevirtual指令
// 当 JVM 通过 invokevirtual 调用方法时, this 引用指向某个类（或其子类）的实例. 因为类的继承层次是固定的,
// 所以虚拟机可以使用一种叫作 vtable（Virtual GetMethod Table）的技术加速方法查找
type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func (self *INVOKE_VIRTUAL) Execute(frame *rtda.Frame) {
	// 拿到当前类、当前常量池、方法符号引用，然后解析符号引用，拿到解析后的类和方法
	currentClass := frame.GetMethod().GetClass()
	constantPool := currentClass.GetConstantPool()
	methodRef := constantPool.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	// 当前方法不能为静态方法
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 取出 this 判断空指针并判断是否为 System.out.println()
	ref := frame.GetOperandStack().GetRefFromTop(resolvedMethod.GetArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	// 断确保protected方法只能被声明该方法的类或子类调用
	if resolvedMethod.IsProtected() &&
		resolvedMethod.GetClass().IsSuperClassOf(currentClass) &&
		resolvedMethod.GetClass().GetPackageName() != currentClass.GetPackageName() &&
		ref.GetClass() != currentClass &&
		!ref.GetClass().IsSubClassOf(currentClass){
		if !(ref.GetClass().IsArrayClass() && resolvedMethod.GetName() == "clone") {
			panic("java.lang.IllegalAccessError")
		}
	}
	// 从对象的类中查找真正要调用的方法。如果找不到方法，或者找到的是抽象方法，则需要抛出AbstractMethodError异常，否则一切正常，调用方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.GetClass(), methodRef.GetName(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func (self *INVOKE_VIRTUAL) String() string {
	return "{type：invoke_virtual; " + self.Index16Instruction.String() + "}"
}