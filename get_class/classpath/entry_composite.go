package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newqCompositeEntry(pathList string) CompositeEntry {
	compositeEntry := []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, form, err := entry.readClass(className)
		if err == nil {
			return data, form, nil
		}
	}
	return nil, nil, errors.New("ClassNotFoundException: " + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for index, item := range self {
		strs[index] = item.String()
	}
	return strings.Join(strs, pathListSeparator)
}