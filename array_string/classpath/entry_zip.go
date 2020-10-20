package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

// 适用于压缩包形式的文件
type ZipEntry struct {
	absPath	string	// 压缩包绝对路径
}

//--------------------------------------------------------------------构造器

// 将路径转为绝对路径存入 self.absPath
func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry {
		absPath: absPath,
	}
}

//--------------------------------------------------------------------功能类方法

// 在给定的压缩包中查找指定的 className, 读取其数据后将二进制数据进行返回
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

//--------------------------------------------------------------------toString

func (self *ZipEntry) String() string {
	return self.absPath
}