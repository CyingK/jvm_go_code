package main

import (
	"fmt"
)

// 启动程序, 解析命令行参数, 创建虚拟机并启动
func main() {
	fmt.Println("" +
		"_________        .__             ____   _________   \n" +
		"\\_   ___ \\___.__.|__| ____    ___\\   \\ /   /     \\  \n" +
		"/    \\  \\<   |  ||  |/    \\  / ___\\   Y   /  \\ /  \\ \n" +
		"\\     \\___\\___  ||  |   |  \\/ /_/  >     /    Y    \\\n" +
		" \\______  / ____||__|___|  /\\___  / \\___/\\____|__  /\n" +
		"        \\/\\/             \\//_____/               \\/ ")
	cmd := parseCmd()
	if cmd.versionFlag {
		fmt.Println("version 1.0.0")
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		newJVM(cmd).start()
	}
}