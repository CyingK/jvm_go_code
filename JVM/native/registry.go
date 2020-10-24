package native

import "jvm_go_code/JVM/rtda"

type NativeMethod func(frame *rtda.Frame)

var registry = map[string]NativeMethod{}

// 类名, 方法名和方法描述符加在一起才能唯一确定一个方法, 所以把它们的组合作为本地方法注册表的键
func Register(className string, methodName string, methodDescriptor string, method NativeMethod) {
	key := className + "~" + methodName + "~" + methodDescriptor
	registry[key] = method
}

// 根据类名, 方法名和方法描述符查找本地方法实现, 如果找不到, 则返回nil
func FindNativeMethod(className string, methodName string, methodDescriptor string) NativeMethod {
	key := className + "~" + methodName + "~" + methodDescriptor
	if method, ok := registry[key]; ok {
		return method
	}
	if methodDescriptor == "()V" {
		if methodName == "initIDs" || methodName == "registerNatives" {
			return emptyNativeMethod
		}
	}
	return nil
}

func emptyNativeMethod(frame *rtda.Frame) {
}