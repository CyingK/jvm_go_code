package heap

// 第一重循环遍历本类及父类, 第二重循环遍历类中的所有方法
func LookupMethodInClass(class *Class, name, descriptor string) *Method {
	for item := class; item != nil; item = item.superClass {
		for _, item := range item.methods {
			if item.name == name && item.descriptor == descriptor {
				return item
			}
		}
	}
	return nil
}

// 第一重循环遍历 interfaces, 第二重循环遍历接口中的所有方法
func lookupMethodInInterfaces(interfaces []*Class, name string, descriptor string) *Method {
	for _, item := range interfaces {
		for _, item := range item.methods {
			if item.name == name && item.descriptor == descriptor {
				return item
			}
		}
		method := lookupMethodInInterfaces(item.interfaces, name, descriptor)
		if method != nil {
			return method
		}
	}
	return nil
}
