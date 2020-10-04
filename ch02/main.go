package main

import (
	"fmt"
	classpath "jvm_go_code/ch02/classpath"
	"strconv"
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
	fmt.Printf("classpath: %s\nclass: %s\nargs: %v\n", cmd.cpOption, cmd.class, cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Println("ClassNotFoundException: " + className)
		return
	}
	for _, item := range classData {
		fmt.Printf("%v", strconv.FormatInt(int64(item), 16))
	}
}
