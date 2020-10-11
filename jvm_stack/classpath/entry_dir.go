package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string	// class 文件的绝对路径
}

/*
 * 字符串化
 */
func (self *DirEntry) String() string {
	return self.absDir
}

/*
 * 根据 className 和 self.absDir 获取读取具体的文件，并返回
 */
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

/*
 * 构造函数
 */
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry {
		absDir: absDir,
	}
}