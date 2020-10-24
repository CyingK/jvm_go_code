package misc

import (
	"jvm_go_code/JVM/native"
	"jvm_go_code/JVM/rtda"
	"jvm_go_code/JVM/rtda/heap"
)

const SUM_MISC_UNSAFE = "sun/misc/Unsafe"

func init() {
	native.Register(SUM_MISC_UNSAFE, "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	native.Register(SUM_MISC_UNSAFE, "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	native.Register(SUM_MISC_UNSAFE, "addressSize", "()I", addressSize)
	native.Register(SUM_MISC_UNSAFE, "objectFieldOffset", "(Ljava/lang/reflect/Field;)J", objectFieldOffset)
	native.Register(SUM_MISC_UNSAFE, "compareAndSwapObject", "(Ljava/lang/Object;JLjava/lang/Object;Ljava/lang/Object;)Z", compareAndSwapObject)
	native.Register(SUM_MISC_UNSAFE, "getIntVolatile", "(Ljava/lang/Object;J)I", getInt)
	native.Register(SUM_MISC_UNSAFE, "compareAndSwapInt", "(Ljava/lang/Object;JII)Z", compareAndSwapInt)
	native.Register(SUM_MISC_UNSAFE, "getObjectVolatile", "(Ljava/lang/Object;J)Ljava/lang/Object;", getObject)
	native.Register(SUM_MISC_UNSAFE, "compareAndSwapLong", "(Ljava/lang/Object;JJJ)Z", compareAndSwapLong)

}

func arrayBaseOffset(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(0)
}

func arrayIndexScale(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(1)
}

func addressSize(frame *rtda.Frame) {
	stack := frame.GetOperandStack()
	stack.PushInt(8)
}

func objectFieldOffset(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	jField := vars.GetRef(1)
	offset := jField.GetIntVar("slot", "I")
	stack := frame.GetOperandStack()
	stack.PushLong(int64(offset))
}

func compareAndSwapObject(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	obj := vars.GetRef(1)
	fields := obj.GetData()
	offset := vars.GetLong(2)
	expected := vars.GetRef(4)
	newVal := vars.GetRef(5)
	if anys, ok := fields.(heap.Slots); ok {
		swapped := _casObj(obj, anys, offset, expected, newVal)
		frame.GetOperandStack().PushBoolean(swapped)
	} else if objs, ok := fields.([]*heap.Object); ok {
		swapped := _casArr(objs, offset, expected, newVal)
		frame.GetOperandStack().PushBoolean(swapped)
	} else {
		panic("todo: compareAndSwapObject!")
	}
}

func _casObj(obj *heap.Object, fields heap.Slots, offset int64, expected, newVal *heap.Object) bool {
	current := fields.GetRef(uint(offset))
	if current == expected {
		fields.SetRef(uint(offset), newVal)
		return true
	} else {
		return false
	}
}
func _casArr(objs []*heap.Object, offset int64, expected, newVal *heap.Object) bool {
	current := objs[offset]
	if current == expected {
		objs[offset] = newVal
		return true
	} else {
		return false
	}
}

func getInt(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	fields := vars.GetRef(1).GetData()
	offset := vars.GetLong(2)
	stack := frame.GetOperandStack()
	if slots, ok := fields.(heap.Slots); ok {
		stack.PushInt(slots.GetInt(uint(offset)))
	} else if shorts, ok := fields.([]int32); ok {
		stack.PushInt(int32(shorts[offset]))
	} else {
		panic("getInt!")
	}
}

func compareAndSwapInt(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	fields := vars.GetRef(1).GetData()
	offset := vars.GetLong(2)
	expected := vars.GetInt(4)
	newVal := vars.GetInt(5)
	if slots, ok := fields.(heap.Slots); ok {
		oldVal := slots.GetInt(uint(offset))
		if oldVal == expected {
			slots.SetInt(uint(offset), newVal)
			frame.GetOperandStack().PushBoolean(true)
		} else {
			frame.GetOperandStack().PushBoolean(false)
		}
	} else if ints, ok := fields.([]int32); ok {
		oldVal := ints[offset]
		if oldVal == expected {
			ints[offset] = newVal
			frame.GetOperandStack().PushBoolean(true)
		} else {
			frame.GetOperandStack().PushBoolean(false)
		}
	} else {
		panic("todo: compareAndSwapInt!")
	}
}

func getObject(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	fields := vars.GetRef(1).GetData()
	offset := vars.GetLong(2)
	if anys, ok := fields.(heap.Slots); ok {
		x := anys.GetRef(uint(offset))
		frame.GetOperandStack().PushRef(x)
	} else if objs, ok := fields.([]*heap.Object); ok {
		x := objs[offset]
		frame.GetOperandStack().PushRef(x)
	} else {
		panic("getObject!")
	}
}

func compareAndSwapLong(frame *rtda.Frame) {
	vars := frame.GetLocalVars()
	fields := vars.GetRef(1).GetData()
	offset := vars.GetLong(2)
	expected := vars.GetLong(4)
	newVal := vars.GetLong(6)
	if slots, ok := fields.(heap.Slots); ok {
		oldVal := slots.GetLong(uint(offset))
		if oldVal == expected {
			slots.SetLong(uint(offset), newVal)
			frame.GetOperandStack().PushBoolean(true)
		} else {
			frame.GetOperandStack().PushBoolean(false)
		}
	} else if longs, ok := fields.([]int64); ok {
		oldVal := longs[offset]
		if oldVal == expected {
			longs[offset] = newVal
			frame.GetOperandStack().PushBoolean(true)
		} else {
			frame.GetOperandStack().PushBoolean(false)
		}
	} else {
		// todo
		panic("todo: compareAndSwapLong!")
	}
}
