package main

import (
	"fmt"
	"jvm_go_code/read_class/classfile"
	"jvm_go_code/read_class/classpath"
	"log"
	"strings"
)

func main() {
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 1.0.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printHelp()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	classPath := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	log.Println("命令行参数......")
	log.Println("类路径：", cmd.cpOption)
	log.Println("待加载的类：", cmd.class)
	log.Println("参数列表：", cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	loadClass(className, classPath)
	//classFile := loadClass(className, classPath)
	//printClassInfo(classFile)
}

func printClassInfo(classFile *classfile.ClassFile) {
	fmt.Printf("version: %v.%v\n", classFile.MajorVersion(), classFile.MinorVersion())
	fmt.Printf("constants count: %v\n", len(classFile.ConstantPool()))
	fmt.Printf("access flags: 0x%x\n", classFile.AccessFlags())
	fmt.Printf("this class: %v\n", classFile.ClassName())
	fmt.Printf("super class: %v\n", classFile.SuperClassName())
	fmt.Printf("interfaces: %v\n", classFile.InterfaceNames())
	fmt.Printf("fields count: %v\n", len(classFile.Fields()))
	for _, field := range classFile.Fields() {
		fmt.Printf(" %s\n", field.Name())
	}
	fmt.Printf("methods count: %v\n", len(classFile.Methods()))
	for _, method := range classFile.Methods() {
		fmt.Printf(" %s\n", method.Name())
	}
}

func loadClass(className string, classPath *classpath.ClassPath) *classfile.ClassFile {
	classData, _, err := classPath.ReadClass(className)
	if err != nil {
		panic(err)
	}
	classFile, err := classfile.Parse(classData)
	if err != nil {
		panic(err)
	}
	return classFile
}
