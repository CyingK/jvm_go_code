package references

import (
	"fmt"
	"jvm_go_code/array_string/instructions/base"
	"jvm_go_code/array_string/rtda"
	"jvm_go_code/array_string/rtda/heap"
)

// 调用静态方法
type INVOKE_STATIC struct {
	base.Index16Instruction
}

// 解析得到方法引用, 如果是非静态方法则直接终端, 如果静态方法所属的类还没有初始化则先初始化, 然后调用 base.InvokeMethod(Frame, GetMethod)
func (self *INVOKE_STATIC) Execute(frame *rtda.Frame) {
	constantPool := frame.GetMethod().GetClass().GetConstantPool()
	methodRef := constantPool.GetConstant(self.Index).(*heap.MethodRef)
	resolvedMethod := methodRef.ResolvedMethod()
	if !resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}
	class := resolvedMethod.GetClass()
	fmt.Printf("invoke_static: 调用静态方法%v.%v%v\n", class.GetName(), resolvedMethod.GetName(), resolvedMethod.Descriptor())
	if !class.GetInitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.GetThread(), class)
		return
	}
	base.InvokeMethod(frame, resolvedMethod)
}

func (self *INVOKE_STATIC) String() string {
	return "{type：invoke_static; " + self.Index16Instruction.String() + "}"
}