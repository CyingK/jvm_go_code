package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// 压缩包类型
type ZipEntry struct {
	absPath	string
}

// 构造函数， 将 path 以绝对路径的形式保存到 absPath 中
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{
		absPath: absPath,
	}
}

/**
 * 首先打开压缩包，遍历每一个文件，找到与 className 同名的文件后，将其打开读取数据
 * 其中任何一步出了错误，都会直接返回

 * zip.OpenReader() 读取 zip 文件
 * reader.file 压缩包中的文件
 * fileData.Open() 打开文件
 * ioutil.ReadAll() 读取文件
*/
func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer reader.Close()
	for _, fileData := range reader.File {
		if fileData.Name == className {
			readClass, err := fileData.Open()
			if err != nil {
				return nil, nil, err
			}
			defer readClass.Close()
			data, err := ioutil.ReadAll(readClass)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("ClassNotFoundException: " + className)
}

// 返回绝对路径
func (self *ZipEntry) String() string {
	return self.absPath
}