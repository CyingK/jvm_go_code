package classpath

import (
	"os"
	"path/filepath"
)

// 类路径
type ClassPath struct {
	bootClassPath Entry 	// 引导类加载路径
	extClassPath  Entry 	// 引导类加载路径
	userClassPath Entry 	// 引导类加载路径
}

//--------------------------------------------------------------------功能类方法

// 分别调用 resolveBootAndExtClasspath, resolveUserClasspath 解析 引导类, 扩展类, 系统类所在的目录路径
func ResolveClassPath(jreOption, cpOption string) *ClassPath {
	classPath := &ClassPath{}
	classPath.resolveBootAndExtClasspath(jreOption)
	classPath.resolveUserClasspath(cpOption)
	return classPath
}

// 根据 jreOption 获取 jreDir 解析引导类所需加载的类路径：{jreDir}/lib/* 和扩展类所需加载的类路径：{jreDir}/lib/ext/*
func (self *ClassPath) resolveBootAndExtClasspath(jreOption string) {
	jreDir := self.resolveJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildcardEntry(jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath = newWildcardEntry(jreExtPath)
}

// 如果显示指定了 jre 路径, 则直接返回, 否则查看当前目录有没有 jre, 有就返回. 否则根据 JAVA_HOME 环境变量来找
func (self *ClassPath) resolveJreDir(jreOption string) string {
	if jreOption != "" && self.exists(jreOption) {
		return jreOption
	}
	if self.exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("JreNotFoundException.")
}

// 如果用户指定了 classpath 则直接使用 classpath 否则就以当前路径作为 classpath
func (self *ClassPath) resolveUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
}

// 分别使用引导类加载器, 扩展类加载器, 系统类加载器去加载指定的 class 文件, 返回读取如来的二进制数据
func (self *ClassPath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	if data, entry, err := self.extClassPath.readClass(className); err == nil {
		return data, entry, err
	}
	return self.userClassPath.readClass(className)
}

// 判断指定路径是否存在
func (self *ClassPath) exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

//--------------------------------------------------------------------toString

func (self *ClassPath) String() string {
	return self.userClassPath.String()
}
