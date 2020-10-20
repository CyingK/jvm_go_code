package heap

type MethodDescriptor struct {
	parameterTypes	[]string	// 参数类型列表
	returnType		string		// 返回值
}

func (self *MethodDescriptor) addParameterType(_type_ string) {
	paramsLen := len(self.parameterTypes)
	if paramsLen == cap(self.parameterTypes) {
		s := make([]string, paramsLen, paramsLen + 4)
		copy(s, self.parameterTypes)
		self.parameterTypes = s
	}
	self.parameterTypes = append(self.parameterTypes, _type_)
}

