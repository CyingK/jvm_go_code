package references

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
)

// 因为私有方法和构造函数不需要动态绑定, 所以invokespecial指令可以加快方法调用速度
type INVOKE_SPECIAL struct {
	base.Index16Instruction
}

func (self *INVOKE_SPECIAL) Execute(frame *rtda.Frame) {
	fmt.Println("invoke_special: 从操作数栈弹出一个引用")
	// 拿到当前类、当前常量池、方法符号引用，然后解析符号引用，拿到解析后的类和方法
	currentClass := frame.GetMethod().GetClass()
	constantPool := currentClass.GetConstantPool()
	methodRef := constantPool.GetConstant(self.Index).(*heap.MethodRef)
	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()
	// 如果当前是 <init> 方法则该方法的声明类必须是 resolvedClass
	if resolvedMethod.GetName() == "<init>" && resolvedMethod.GetClass() != resolvedClass {
		panic("java.lang.NoSuchMethodError")
	}
	// 当前方法不能为静态方法
	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	// 取出 this 判断空指针
	ref := frame.GetOperandStack().GetRefFromTop(resolvedMethod.GetArgSlotCount() - 1)
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
	// 断确保protected方法只能被声明该方法的类或子类调用
	if resolvedMethod.IsProtected() &&
		resolvedMethod.GetClass().IsSuperClassOf(currentClass) &&
		resolvedMethod.GetClass().GetPackageName() != currentClass.GetPackageName() &&
		ref.GetClass() != currentClass &&
		!ref.GetClass().IsSubClassOf(currentClass) {
		panic("java.lang.IllegalAccessError")
	}
	// 如果当前类被标记为 ACC_SUPER, 且 resolvedClass 是 currentClass 的子类, 但 resolvedMethod 不是构造函数
	// 需要一个额外的过程查找最终要调用的方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		resolvedClass.IsSuperClassOf(currentClass) &&
		resolvedMethod.GetName() != "<init>"{
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.GetSuperClass(), methodRef.Name(), methodRef.Descriptor())
	}
	// 如果查找失败或查找到的是抽象方法报错
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func (self *INVOKE_SPECIAL) String() string {
	return "{type：invoke_special; " + self.Index16Instruction.String() + "}"
}