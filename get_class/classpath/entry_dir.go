package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 基本类型
type DirEntry struct {
	absDir string
}

// 构造函数, 将 path 以绝对路径的形式保存到 absDir 中
func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{
		absDir: absDir,
	}
}

/**
 * 读取 class 文件, 根据 class 名和 absDir 获取读取具体的文件，并返回

 * filepath.Join() 拼接路径
 * ioutil.ReadFile() 读取文件，返回二进制数据
*/
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

// 返回绝对路径
func (self *DirEntry) String() string {
	return self.absDir
}