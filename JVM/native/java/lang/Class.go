package lang

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
	"strings"
)

const JAVA_LANG_CLASS = "java/lang/Class"

func init() {
	native.Register(JAVA_LANG_CLASS, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.Register(JAVA_LANG_CLASS, "getName0", "()Ljava/lang/String;", getName0)
	native.Register(JAVA_LANG_CLASS, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
	native.Register(JAVA_LANG_CLASS, "isInterface", "()Z", isInterface)
	native.Register(JAVA_LANG_CLASS, "isPrimitive", "()Z", isPrimitive)
	native.Register(JAVA_LANG_CLASS, "getDeclaredFields0", "(Z)[Ljava/lang/reflect/Field;", getDeclaredFields0)
	native.Register(JAVA_LANG_CLASS, "forName0", "(Ljava/lang/String;ZLjava/lang/ClassLoader;Ljava/lang/Class;)Ljava/lang/Class;", forName0)
	native.Register(JAVA_LANG_CLASS, "getDeclaredConstructors0", "(Z)[Ljava/lang/reflect/Constructor;", getDeclaredConstructors0)
	native.Register(JAVA_LANG_CLASS, "getModifiers", "()I", getModifiers)
	native.Register(JAVA_LANG_CLASS, "getSuperclass", "()Ljava/lang/Class;", getSuperclass)
	native.Register(JAVA_LANG_CLASS, "getInterfaces0", "()[Ljava/lang/Class;", getInterfaces0)
	native.Register(JAVA_LANG_CLASS, "isArray", "()Z", isArray)
	native.Register(JAVA_LANG_CLASS, "getDeclaredMethods0", "(Z)[Ljava/lang/reflect/Method;", getDeclaredMethods0)
	native.Register(JAVA_LANG_CLASS, "getComponentType", "()Ljava/lang/Class;", getComponentType)
	native.Register(JAVA_LANG_CLASS, "isAssignableFrom", "(Ljava/lang/Class;)Z", isAssignableFrom)
}

// (Ljava/lang/Class;)Z
func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.GetOperandStack().PushBoolean(false)
}

// 从局部变量表中拿到 this 引用, 这是一个类对象引用, 通过 GetExtra() 方法可以获得与之对应的
// Class 结构体指针. 然后拿到类名, 转成Java字符串并推入操作数栈顶
func getName0(frame *rtda.Frame) {
	this := frame.GetLocalVars().GetThis()
	class := this.GetExtra().(*heap.Class)
	name := class.GetJavaName()
	nameObject := heap.ToJavaString(class.GetClassLoader(), name)
	frame.GetOperandStack().PushRef(nameObject)
}

// 静态方法. 先从局部变量表中拿到类名, 这是个Java字符串, 需要把它转成Go字符串. 基本类型的类已经加
// 载到了方法区中, 直接调用类加载器的 LoadClass() 方法获取即可. 最后, 把类对象引用推入操作数栈顶
func getPrimitiveClass(frame *rtda.Frame) {
	nameObject := frame.GetLocalVars().GetRef(0)
	name := heap.ToGoString(nameObject)
	classLoader := frame.GetMethod().GetClass().GetClassLoader()
	class := classLoader.LoadClass(name).GetJavaClass()
	frame.GetOperandStack().PushRef(class)
}

func isInterface(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	this := vars.GetThis()
	class := this.GetExtra().(*heap.Class)

	stack := frame.GetOperandStack()
	stack.PushBoolean(class.IsInterface())
}

func isPrimitive(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	this := vars.GetThis()
	class := this.GetExtra().(*heap.Class)

	stack := frame.GetOperandStack()
	stack.PushBoolean(class.IsPrimitive())
}

func forName0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	jName := vars.GetRef(0)
	initialize := vars.GetBoolean(1)
	goName := heap.ToGoString(jName)
	goName = strings.Replace(goName, ".", "/", -1)
	goClass := frame.GetMethod().GetClass().GetClassLoader().LoadClass(goName)
	jClass := goClass.GetJavaClass()
	if initialize && !goClass.GetInitStarted() {
		thread := frame.GetThread()
		frame.SetNextPC(thread.GetPC())
		base.InitClass(thread, goClass)
	} else {
		stack := frame.GetOperandStack()
		stack.PushRef(jClass)
	}
}

func getModifiers(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	this := vars.GetThis()
	class := this.GetExtra().(*heap.Class)
	modifiers := class.GetAccessFlags()
	stack := frame.GetOperandStack()
	stack.PushInt(int32(modifiers))
}

func getSuperclass(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	this := vars.GetThis()
	class := this.GetExtra().(*heap.Class)
	superClass := class.GetSuperClass()
	stack := frame.GetOperandStack()
	if superClass != nil {
		stack.PushRef(superClass.GetJavaClass())
	} else {
		stack.PushRef(nil)
	}
}

func getInterfaces0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	this := vars.GetThis()
	class := this.GetExtra().(*heap.Class)
	interfaces := class.GetInterfaces()
	classArr := toClassArr(class.GetClassLoader(), interfaces)
	stack := frame.GetOperandStack()
	stack.PushRef(classArr)
}

func isArray(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	this := vars.GetThis()
	class := this.GetExtra().(*heap.Class)
	stack := frame.GetOperandStack()
	stack.PushBoolean(class.IsArrayClass())
}

func getComponentType(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	this := vars.GetThis()
	class := this.GetExtra().(*heap.Class)
	componentClass := class.GetComponentClass()
	componentClassObj := componentClass.GetJavaClass()
	stack := frame.GetOperandStack()
	stack.PushRef(componentClassObj)
}

func isAssignableFrom(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	this := vars.GetThis()
	cls := vars.GetRef(1)
	thisClass := this.GetExtra().(*heap.Class)
	clsClass := cls.GetExtra().(*heap.Class)
	ok := thisClass.IsAssignableFrom(clsClass)
	stack := frame.GetOperandStack()
	stack.PushBoolean(ok)
}
