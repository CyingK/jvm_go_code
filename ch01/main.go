package main

import "fmt"

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
	fmt.Printf("classpath: %s\nclass: %s\nargs: %v\n",cmd.cpOption, cmd.class, cmd.args)
}
