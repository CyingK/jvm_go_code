package heap

type Object struct {
	class *Class
	data  interface{}
}

func newObject(class *Class) *Object {
	return &Object {
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
}

func (self *Object) GetClass() *Class {
	return self.class
}

func (self *Object) GetFields() Slots {
	return self.data.(Slots)
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) String() string {
	return "nil"
}

//--------------------------------------------------------------------Getters

func (self *Object) GetBytes() []int8 {
	return self.data.([]int8)
}

func (self *Object) GetShorts() []int16 {
	return self.data.([]int16)
}

func (self *Object) GetInts() []int32 {
	return self.data.([]int32)
}

func (self *Object) GetLongs() []int64 {
	return self.data.([]int64)
}

func (self *Object) GetChars() []uint16 {
	return self.data.([]uint16)
}

func (self *Object) GetFloats() []float32 {
	return self.data.([]float32)
}

func (self *Object) GetDoubles() []float64 {
	return self.data.([]float64)
}

func (self *Object) GetRefs() []*Object {
	return self.data.([]*Object)
}

func (self *Object) GetArrayLength() int32 {
	switch self.data.(type) {
	case []int8:
		return int32(len(self.data.([]int8)))
	case []int16:
		return int32(len(self.data.([]int16)))
	case []int32:
		return int32(len(self.data.([]int32)))
	case []int64:
		return int32(len(self.data.([]int64)))
	case []uint16:
		return int32(len(self.data.([]uint16)))
	case []float32:
		return int32(len(self.data.([]float32)))
	case []float64:
		return int32(len(self.data.([]float64)))
	case []*Object:
		return int32(len(self.data.([]*Object)))
	default:
		panic("Not array!")
	}
}

//--------------------------------------------------------------------Setters

// 直接给对象的引用类型实例变量赋值
func (self *Object) SetRefVar(name string, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) GetRefVar(name string, descriptor string) *Object {
	field := self.GetClass().getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}