package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 文件入口, 适用于指定单一文件路径的情况
type DirEntry struct {
	absDir string	// class 文件的绝对路径
}

//--------------------------------------------------------------------构造器

// 创建 DirEntry
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry {
		absDir: absDir,
	}
}

//--------------------------------------------------------------------功能类方法

// 根据 className 和 self.absDir 获取并读取具体的文件, 将读出的二进制数据返回
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

//--------------------------------------------------------------------toString

func (self *DirEntry) String() string {
	return self.absDir
}