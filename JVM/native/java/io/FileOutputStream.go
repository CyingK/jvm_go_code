package io

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"os"
	"unsafe"
)

const JAVA_LANG_FILEOUTPUTSTREAM = "java/io/FileOutputStream"

func init() {
	native.Register(JAVA_LANG_FILEOUTPUTSTREAM, "writeBytes", "([BIIZ)V", writeBytes)
}

func writeBytes(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	b := vars.GetRef(1)
	off := vars.GetInt(2)
	len := vars.GetInt(3)
	jBytes := b.GetData().([]int8)
	go_bytes := castInt8sToUint8s(jBytes)
	go_bytes = go_bytes[off: off + len]
	os.Stdout.Write(go_bytes)
}

func castInt8sToUint8s(java_bytes []int8) (go_bytes []byte) {
	pointer := unsafe.Pointer(&java_bytes)
	go_bytes = *((*[]byte)(pointer))
	return
}
