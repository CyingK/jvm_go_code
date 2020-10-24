package references

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
	"reflect"
)

// Throw exception or error
type ATHROW struct{ base.NoOperandsInstruction }

// 先从操作数栈中弹出异常对象引用, 如果该引用是null, 则抛出 NullPointerException 异常
// 否则看是否可以找到并跳转到异常处理代码
func (self *ATHROW) Execute(frame *rtda.Frame) {
	ex := frame.GetOperandStack().PopRef()
	if ex == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.GetThread()
	if !findAndGotoExceptionHandler(thread, ex) {
		handleUncaughtException(thread, ex)
	}
}

// 从当前帧开始, 遍历Java虚拟机栈, 查找方法的异常处理表. 假设遍历到帧, 如果在F对应的方法中找不到异常处理项, 则把F
// 弹出, 继续遍历. 反之如果找到了异常处理项, 在跳转到异常处理代码之前, 要先把F的操作数栈清空, 然后把异常对象引用推入栈
// 顶如果遍历完Java虚拟机栈还是找不到异常处理代码, 则 handleUncaughtException() 函数打印出Java虚拟机栈信息
func findAndGotoExceptionHandler(thread *rtda.Thread, ex *heap.Object) bool {
	for {
		frame := thread.GetCurrentFrame()
		pc := frame.GetNextPC() - 1

		handlerPC := frame.GetMethod().FindExceptionHandler(ex.GetClass(), pc)
		if handlerPC > 0 {
			stack := frame.GetOperandStack()
			stack.Clear()
			stack.PushRef(ex)
			frame.SetNextPC(handlerPC)
			return true
		}

		thread.PopFrame()
		if thread.IsStackEmpty() {
			break
		}
	}
	return false
}

// handleUncaughtException() 函数把Java虚拟机栈清空, 然后打印出异常信息
func handleUncaughtException(thread *rtda.Thread, ex *heap.Object) {
	thread.ClearStack()

	jMsg := ex.GetRefVar("detailMessage", "Ljava/lang/String;")
	goMsg := heap.ToGoString(jMsg)
	println(ex.GetClass().GetJavaName() + ": " + goMsg)

	stes := reflect.ValueOf(ex.GetExtra())
	for i := 0; i < stes.Len(); i++ {
		ste := stes.Index(i).Interface().(interface {
			String() string
		})
		println("\tat " + ste.String())
	}
}
