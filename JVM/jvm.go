package main

import (
	"fmt"
	"jvm_go_code/JVM/classpath"
	"jvm_go_code/JVM/instructions/base"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
	"strings"
)

// Java Virtual Machine
type JVM struct {
	cmd				*Cmd					// 命令行参数
	classLoader     *heap.ClassLoader		// 类加载器
	mainThread		*rtda.Thread			// 主线程
}

// 执行 main 方法
func (self *JVM) executeMain() {
	className := strings.Replace(self.cmd.class, ".", "/", -1)
	mainClass := self.classLoader.LoadClass(className)
	mainMethod := mainClass.GetMainMethod()
	if mainMethod == nil {
		fmt.Printf("Main method not found in class %s\n", self.cmd.class)
		return
	}
	argsArr := func () *heap.Object {
		stringClass := self.classLoader.LoadClass("java/lang/String")
		argsLen := uint(len(self.cmd.args))
		argsArr := stringClass.GetArrayClass().NewArray(argsLen)
		java_args := argsArr.GetRefs()
		for index, arg := range self.cmd.args {
			java_args[index] = heap.ToJavaString(self.classLoader, arg)
		}
		return argsArr
	}()
	frame := self.mainThread.NewFrame(mainMethod)
	frame.GetLocalVars().SetRef(0, argsArr)
	self.mainThread.PushFrame(frame)
	interpret(self.mainThread, self.cmd.verboseInstanceFlag)
}

// 初始化 Java 虚拟机：加载 sun/misc/VM 类, 初始化类并执行解释器
func (self *JVM) initializeVM() {
	vmClass := self.classLoader.LoadClass("sun/misc/VM")
	base.InitClass(self.mainThread, vmClass)
	interpret(self.mainThread, self.cmd.verboseInstanceFlag)
}

// 启动 Java 虚拟机
func (self *JVM) start() {
	self.initializeVM()
	self.executeMain()
}

// 创建 Java 虚拟机：解析类路径, 创建类加载器, 返回创建的 Java 虚拟机
func newJVM(cmd *Cmd) *JVM {
	classPath := classpath.ResolveClassPath(cmd.XjreOption, cmd.cpOption)
	classLoader := heap.NewClassLoader(classPath, cmd.verboseClassFlag)
	return &JVM {
		cmd: 			cmd,
		classLoader: 	classLoader,
		mainThread: 	rtda.NewThread(),
	}
}