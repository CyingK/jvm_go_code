package classpath

import (
	"errors"
	"strings"
)

// 组合入口, 适用于指定了多个 classPath 的情况
type CompositeEntry []Entry

//--------------------------------------------------------------------构造器

// 将 path 以间隔符为标志拆分，对拆分后的每一个路径分别创建 Entry
func newCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

//--------------------------------------------------------------------功能类方法

// 遍历每一个 Entry 元素的内容查找 className 并读取
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, form, err := entry.readClass(className)
		if err == nil {
			return data, form, nil
		}
	}
	return nil, nil, errors.New("ClassNotFoundException: " + className)
}


//--------------------------------------------------------------------toString

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for index, item := range self {
		strs[index] = item.String()
	}
	return strings.Join(strs, pathListSeparator)
}