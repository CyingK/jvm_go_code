package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

//--------------------------------------------------------------------构造器

// 根据 baseDir 查找该文件夹中的所有 jar 文件, 封装成一个 []Entry
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path) - 1]
	compositeEntry := []Entry{}
	filepath.Walk(baseDir, func (path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}
		return nil
	})
	return compositeEntry
}
