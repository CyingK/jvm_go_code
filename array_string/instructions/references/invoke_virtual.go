package references

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
)

// 使用 super 关键字调用超类中的方法不能使用invokevirtual指令
// 当 JVM 通过 invokevirtual 调用方法时, this 引用指向某个类（或其子类）的实例. 因为类的继承层次是固定的,
// 所以虚拟机可以使用一种叫作 vtable（Virtual GetMethod Table）的技术加速方法查找
type INVOKE_VIRTUAL struct {
	base.Index16Instruction
}

func util_print(stack *rtda.OperandStack, descriptor string) {
	switch descriptor {
	case "(Z)V":
		fmt.Printf("\nConsole >  %v\n", stack.PopInt() != 0)
	case "(C)V":
		fmt.Printf("\nConsole >  %c\n", stack.PopInt())
	case "(I)V", "(B)V", "(S)V":
		fmt.Printf("\nConsole >  %v\n", stack.PopInt())
	case "(F)V":
		fmt.Printf("\nConsole >  %v\n", stack.PopFloat())
	case "(J)V":
		fmt.Printf("vConsole >  %v\n", stack.PopLong())
	case "(D)V":
		fmt.Printf("\nConsole >  %v\n", stack.PopDouble())
	case "(Ljava/lang/String;)V":
		java_string := stack.PopRef()
		golang_string := heap.JStringToGoString(java_string)
		fmt.Printf("Console >  %v\n", golang_string)
	default:
		panic("println: " + descriptor)
	}
	stack.PopRef()
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
		if methodRef.Name() == "println" {
			fmt.Printf("invoke_virtual: 调用静态方法%v.%v%v\n", "java/io/PrintStream", resolvedMethod.GetName(), resolvedMethod.Descriptor())
			util_print(frame.GetOperandStack(), methodRef.Descriptor())
			return
		}
		panic("java.lang.NullPointerException")
	}
	// 断确保protected方法只能被声明该方法的类或子类调用
	if resolvedMethod.IsProtected() &&
		resolvedMethod.GetClass().IsSuperClassOf(currentClass) &&
		resolvedMethod.GetClass().GetPackageName() != currentClass.GetPackageName() &&
		ref.GetClass() != currentClass &&
		!ref.GetClass().IsSubClassOf(currentClass){
		panic("java.lang.IllegalAccessError")
	}
	// 从对象的类中查找真正要调用的方法。如果找不到方法，或者找到的是抽象方法，则需要抛出AbstractMethodError异常，否则一切正常，调用方法
	methodToBeInvoked := heap.LookupMethodInClass(ref.GetClass(), methodRef.Name(), methodRef.Descriptor())
	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}
	base.InvokeMethod(frame, methodToBeInvoked)
}

func (self *INVOKE_VIRTUAL) String() string {
	return "{type：invoke_virtual; " + self.Index16Instruction.String() + "}"
}