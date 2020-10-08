package main

import (
	"fmt"
	classpath "jvm_go_code/ch02/classpath"
	"log"
	//"strconv"
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
	log.Println("命令行参数......")
	log.Println("类路径：", cmd.cpOption)
	log.Println("待加载的类：", cmd.class)
	log.Println("参数列表：", cmd.args)
	className := strings.Replace(cmd.class, ".", "/", -1)
	classData, _, err := cp.ReadClass(className)
	if err != nil {
		fmt.Println("ClassNotFoundException: " + className)
		return
	}
	var count int64 = 0
	for _, item := range classData {
		fmt.Printf("%02X ", item)
		if count++; count % 64 == 0 {
			println()
		}
	}
}
