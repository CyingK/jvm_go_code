package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/*
 * 根据 path 获得 baseDir
 * 遍历 baseDir 下的每一项
 * 遍历时出错则返回错误
 * 遍历到子文件夹则直接跳过
 * 遍历到 jar 包则新建 ZipEntry
 */
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
