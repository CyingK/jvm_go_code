package io

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
	"os"
	"path/filepath"
)

const JAVA_IO_UNIXFILESYSTEM = "java/io/UnixFileSystem"


func init() {
	native.Register(JAVA_IO_UNIXFILESYSTEM, "canonicalize0", "(Ljava/lang/String;)Ljava/lang/String;", canonicalize0)
	native.Register(JAVA_IO_UNIXFILESYSTEM, "getBooleanAttributes0", "(Ljava/io/File;)I", getBooleanAttributes0)

}

func canonicalize0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	path := vars.GetRef(1)
	goPath := heap.ToGoString(path)
	goPath2 := filepath.Clean(goPath)
	if goPath2 != goPath {
		path = heap.ToJavaString(frame.GetMethod().GetClass().GetClassLoader(), goPath2)
	}

	stack := frame.GetOperandStack()
	stack.PushRef(path)
}

func getBooleanAttributes0(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	f := vars.GetRef(1)
	path := _getPath(f)
	attributes0 := 0
	if _exists(path) {
		attributes0 |= 0x01
	}
	if _isDir(path) {
		attributes0 |= 0x04
	}
	stack := frame.GetOperandStack()
	stack.PushInt(int32(attributes0))
}

func _getPath(fileObj *heap.Object) string {
	pathStr := fileObj.GetRefVar("path", "Ljava/lang/String;")
	return heap.ToGoString(pathStr)
}

func _exists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func _isDir(path string) bool {
	fileInfo, err := os.Stat(path)
	if err == nil {
		return fileInfo.IsDir()
	}
	return false
}
