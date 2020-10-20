package main

import (
	"fmt"
	"jvm_go_code/array_string/classpath"
	"jvm_go_code/array_string/rtda/heap"
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

// 准备类路径
// 创建类加载器
// 路径替换
// 加载主类
// 捕捉 main 方法
// 开始执行 main 方法
func startJVM(cmd *Cmd) {
	classPath := classpath.ResolveClassPath(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(classPath, cmd.verboseClassFlag)
	className := strings.Replace(cmd.class, ".", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod, cmd.verboseInstanceFlag, cmd.args)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}


//--------------------------------------------------------------------功能类方法
//--------------------------------------------------------------------toString
//--------------------------------------------------------------------判断类方法
//--------------------------------------------------------------------构造器
//--------------------------------------------------------------------Getters