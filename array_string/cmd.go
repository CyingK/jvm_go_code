package main

import (
	"flag"
	"fmt"
	"os"
)

type Cmd struct {
	helpFlag 			bool
	versionFlag			bool
	verboseClassFlag	bool
	verboseInstanceFlag	bool
	cpOption			string
	XjreOption			string
	class				string
	args				[]string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}
	flag.Usage = printHelp
	flag.BoolVar(&cmd.helpFlag, "help", false, "Print Help Message.")
	flag.BoolVar(&cmd.helpFlag, "?", false, "Print Help Message.")
	flag.BoolVar(&cmd.versionFlag, "version", false, "Print Version And Exit.")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseClassFlag, "verbose:class", false, "enable verbose output")
	flag.BoolVar(&cmd.verboseInstanceFlag, "verbose:inst", false, "enable verbose output")
	flag.StringVar(&cmd.cpOption, "classpath", "", "Path Of GetClass")
	flag.StringVar(&cmd.cpOption, "cp", "", "Path Of GetClass")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "Root Of Jre")
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