package lang

import (
	"fmt"
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

const JAVA_LANG_THROWABLE = "java/lang/Throwable"

type StackTraceElement struct {
	fileName   string	// 文件名
	className  string	// 声明方法的类名
	methodName string	// 方法名
	lineNumber int		// 帧正在执行的行号
}

func (self *StackTraceElement) String() string {
	return fmt.Sprintf("%s.%s(%s:%d)",
		self.className, self.methodName, self.fileName, self.lineNumber)
}

func init() {
	native.Register(JAVA_LANG_THROWABLE, "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {
	this := frame.GetLocalVars().GetThis()
	frame.GetOperandStack().PushRef(this)

	stes := createStackTraceElements(this, frame.GetThread())
	this.SetExtra(stes)
}

func createStackTraceElements(tObj *heap.Object, thread *rtda.Thread) []*StackTraceElement {
	skip := distanceToObject(tObj.GetClass()) + 2
	frames := thread.GetFrames()[skip:]
	stes := make([]*StackTraceElement, len(frames))
	for i, frame := range frames {
		stes[i] = createStackTraceElement(frame)
	}
	return stes
}

func distanceToObject(class *heap.Class) int {
	distance := 0
	for c := class.GetSuperClass(); c != nil; c = c.GetSuperClass() {
		distance++
	}
	return distance
}

func createStackTraceElement(frame *rtda.Frame) *StackTraceElement {
	method := frame.GetMethod()
	class := method.GetClass()
	return &StackTraceElement{
		fileName:   class.GetSourceFile(),
		className:  class.GetJavaName(),
		methodName: method.GetName(),
		lineNumber: method.GetLineNumber(frame.GetNextPC() - 1),
	}
}
