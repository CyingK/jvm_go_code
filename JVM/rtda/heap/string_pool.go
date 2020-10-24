package heap

import (
	"unicode/utf16"
)



var internedStrings = map[string]*Object{}

// 根据Go字符串返回相应的Java字符串实例。如果Java字符串已经在池中，直接返回即可，否则先把Go字符串（UTF8格式）转换成Java字符数组（UTF16格式）
// 然后创建一个Java字符串实例，把它的value变量设置成刚刚转换而来的字符数组，最后把Java字符串放入池中
func ToJavaString(classLoader *ClassLoader, key string) *Object {
	if internedStr, ok := internedStrings[key]; ok {
		return internedStr
	}
	chars := stringToUTF16(key)
	java_chars := &Object{
		class: 	classLoader.LoadClass("[C"),
		data: 	chars,
	}
	java_string := classLoader.LoadClass("java/lang/String").NewObject()
	java_string.SetRefVar("value", "[C", java_chars)
	internedStrings[key] = java_string
	return java_string
}

func ToGoString(java_string *Object) string {
	charArray := java_string.GetRefVar("value", "[C")
	return utf16ToString(charArray.GetChars())
}


// 强制转成UTF32
func stringToUTF16(key string) []uint16 {
	runes := []rune(key)
	return utf16.Encode(runes)
}

func utf16ToString(s []uint16) string {
	runes := utf16.Decode(s)
	return string(runes)
}

func InternString(java_string *Object) *Object {
	go_string := ToGoString(java_string)
	if interned, ok := internedStrings[go_string]; ok {
		return interned
	}
	internedStrings[go_string] = java_string
	return java_string
}