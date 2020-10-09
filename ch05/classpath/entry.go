package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)		// 从 class 文件中读取数据
	String() string										// 字符串化
}

/*
 * 路径中有分隔符，使用 CompositeEntry
 * 路径最后是 *，使用 WildcardEntry
 * 路径是压缩包，使用 ZipEntry
 * 普通路径使用 DirEntry
 */
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
	   strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}