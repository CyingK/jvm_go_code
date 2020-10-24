package lang

import (
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
	"runtime"
	"time"
)

const JAVA_LANG_SYSTEM = "java/lang/System"

func init() {
	native.Register(JAVA_LANG_SYSTEM, "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arraycopy)
	native.Register(JAVA_LANG_SYSTEM, "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	native.Register(JAVA_LANG_SYSTEM, "setIn0", "(Ljava/io/InputStream;)V", setIn0)
	native.Register(JAVA_LANG_SYSTEM, "setOut0", "(Ljava/io/PrintStream;)V", setOut0)
	native.Register(JAVA_LANG_SYSTEM, "setErr0", "(Ljava/io/PrintStream;)V", setErr0)
	native.Register(JAVA_LANG_SYSTEM, "currentTimeMillis", "()J", currentTimeMillis)
}

func arraycopy(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)
	// 源数组和目标数组都不能是 null
	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}
	// 源数组和目标数组必须兼容才能拷贝
	if !checkArrayCopy(src, dest) {
		panic("java.lang.ArrayStoreException")
	}
	// 检查srcPos, destPos和length参数
	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos + length > src.GetArrayLength() ||
		destPos + length > dest.GetArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}
	heap.ArrayCopy(src, dest, srcPos, destPos, length)
}

// 首先确保 src 和 dest 都是数组, 然后检查数组类型. 如果两者都是引用数组, 则可以拷贝, 否则两者必须是
// 相同类型的基本类型数组
func checkArrayCopy(src *heap.Object, dest *heap.Object) bool {
	srcClass := src.GetClass()
	destClass := dest.GetClass()
	if !srcClass.IsArrayClass() || !destClass.IsArrayClass() {
		return false
	}
	if srcClass.GetComponentClass().IsPrimitive() ||
		destClass.GetComponentClass().IsPrimitive() {
		return srcClass == destClass
	}
	return true
}

func initProperties(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	props := vars.GetRef(0)

	stack := frame.GetOperandStack()
	stack.PushRef(props)

	// public synchronized Object setProperty(String key, String value)
	setPropMethod := props.GetClass().GetInstanceMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")
	thread := frame.GetThread()
	for key, val := range _sysProps() {
		jKey := heap.ToJavaString(frame.GetMethod().GetClass().GetClassLoader(), key)
		jVal := heap.ToJavaString(frame.GetMethod().GetClass().GetClassLoader(), val)
		ops := rtda.NewOperandStack(3)
		ops.PushRef(props)
		ops.PushRef(jKey)
		ops.PushRef(jVal)
		shimFrame := rtda.NewShimFrame(thread, ops)
		thread.PushFrame(shimFrame)
		base.InvokeMethod(shimFrame, setPropMethod)
	}
}

func _sysProps() map[string]string {
	return map[string]string{
		"java.version":         "1.8.0",
		"java.vendor":          "jvm.go",
		"java.vendor.url":      "",
		"java.home":            "todo",
		"java.class.version":   "52.0",
		"java.class.path":      "todo",
		"java.awt.graphicsenv": "sun.awt.CGraphicsEnvironment",
		"os.name":              runtime.GOOS,
		"os.arch":              runtime.GOARCH,
		"os.version":           "",
		"file.separator":       "/",
		"path.separator":       ":",
		"line.separator":       "\n",
		"user.name":            "",
		"user.home":            "",
		"user.dir":             ".",
		"user.country":         "CN",
		"file.encoding":        "UTF-8",
		"sun.stdout.encoding":  "UTF-8",
		"sun.stderr.encoding":  "UTF-8",
	}
}

func setIn0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	in := vars.GetRef(0)

	sysClass := frame.GetMethod().GetClass()
	sysClass.SetStaticRefVar("in", "Ljava/io/InputStream;", in)
}

func setOut0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	out := vars.GetRef(0)

	sysClass := frame.GetMethod().GetClass()
	sysClass.SetStaticRefVar("out", "Ljava/io/PrintStream;", out)
}

func setErr0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	err := vars.GetRef(0)

	sysClass := frame.GetMethod().GetClass()
	sysClass.SetStaticRefVar("err", "Ljava/io/PrintStream;", err)
}

func currentTimeMillis(frame *rtda.Frame) {
	millis := time.Now().UnixNano() / int64(time.Millisecond)
	stack := frame.GetOperandStack()
	stack.PushLong(millis)
}
