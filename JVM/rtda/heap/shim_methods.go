package heap

var (
	_shimClass = &Class{name: "~shim"}
	_returnCode = []byte{0xB1}
	_athrowCode = []byte{0xBF}
	_returnMethod = &Method {
		ClassMember: ClassMember {
			accessFlags: 	ACC_STATIC,
			name: 			"<return>",
			class: 			_shimClass,
		},
		code: _returnCode,
	}
	_athrowMethod = &Method {
		ClassMember: ClassMember {
			accessFlags: 	ACC_STATIC,
			name: 			"<athrow>",
			class: 			_shimClass,
		},
		code: _athrowCode,
	}
)

func ShimReturnMethod() *Method {
	return _returnMethod
}

func ShimAthrowMethod() *Method {
	return _athrowMethod
}