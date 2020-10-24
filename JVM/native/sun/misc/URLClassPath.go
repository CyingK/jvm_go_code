package misc

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
)

func init() {
	native.Register("sun/misc/URLClassPath", "getLookupCacheURLs", "(Ljava/lang/ClassLoader;)[Ljava/net/URL;", getLookupCacheURLs)
}

func getLookupCacheURLs(frame *rtda.Frame) {
	frame.GetOperandStack().PushRef(nil)
}