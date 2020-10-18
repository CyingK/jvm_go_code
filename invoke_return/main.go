package main

import (
	"fmt"
	"jvm_go_code/invoke_return/classpath"
	"jvm_go_code/invoke_return/rtda/heap"
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
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(cp)
	className := strings.Replace(cmd.class, ",", "/", -1)
	mainClass := classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		fmt.Printf("Main method not found in class %s\n", cmd.class)
	}
}
