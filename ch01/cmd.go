package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag 		bool
	versionFlag		bool
	cpOption		string
	class			string
	args			[]string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printHelp
	flag.BoolVar(&cmd.helpFlag, "help", false, "Print Help Message.")
	flag.BoolVar(&cmd.helpFlag, "?", false, "Print Help Message.")
	flag.BoolVar(&cmd.versionFlag, "version", false, "Print Version And Exit.")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.Parse()
	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}
	return cmd
}

func printHelp() {
	fmt.Printf("Usage: %s [-options] class [args...]\n", os.Args[0])
}