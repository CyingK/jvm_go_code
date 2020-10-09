package classpath

import (
	"log"
	"os"
	"path/filepath"
)

type ClassPath struct {
	bootClassPath Entry // 引导类加载器
	extClassPath  Entry // 扩展类加载器
	userClassPath Entry // 系统类加载器
}

/*
 * 分别使用引导类加载器、扩展类加载器、系统类加载器去加载指定的 class 文件
 */
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

/*
 * 字符串化
 */
func (self *ClassPath) String() string {
	return self.userClassPath.String()
}

/*
 * 根据 jreOption 获取 jreDir
 * 解析引导类所需加载的类路径：{jreDir}/lib/* 和扩展类所需加载的类路径：{jreDir}/lib/ext/*
 */
func (self *ClassPath) parseBootAndExtClasspath(jreOption string) {
	jreDir := GetJreDir(jreOption)
	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClassPath = newWildcardEntry(jreLibPath)
	log.Println("启动类加载器加载路径：", jreLibPath)
	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClassPath = newWildcardEntry(jreExtPath)
	log.Println("扩展类加载器加载路径：", jreExtPath)
}

/*
 * 如果用户指定了 classpath 则直接使用 classpath
 * 否则就以当前路径作为 classpath
 */
func (self *ClassPath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClassPath = newEntry(cpOption)
	log.Println("系统类加载器加载路径：", cpOption)
}

/*
 * 分别解析引导类、扩展类、系统类所在路径
 */
func Parse(jreOption, cpOption string) *ClassPath {
	log.Println("解析加载路径......")
	classPath := &ClassPath{}
	classPath.parseBootAndExtClasspath(jreOption)
	classPath.parseUserClasspath(cpOption)
	return classPath
}

/*
 * 判断指定路径是否存在
 */
func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

/*
 * 获取 jre 路径
 * 如果显示指定了 jre 路径，则直接返回
 * 否则查看当前目录有没有 jre，有就返回
 * 最后根据 JAVA_HOME 环境变量来找
 */
func GetJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("JreNotFoundException.")
}