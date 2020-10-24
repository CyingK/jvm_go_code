package heap

type Object struct {
	class 	*Class			// 对象类
	data  	interface{}		// 对象数据
	extra	interface{}		// 记录Object结构体实例的额外信息
}

func newObject(class *Class) *Object {
	return &Object {
		class: class,
		data:  newSlots(class.instanceSlotCount),
	}
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

func (self *Object) GetClass() *Class {
	return self.class
}

func (self *Object) GetFields() Slots {
	return self.data.(Slots)
}

func (self *Object) GetExtra() interface{} {
	return self.extra
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

func (self *Object) GetRefVar(name string, descriptor string) *Object {
	field := self.GetClass().getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetRef(field.slotId)
}

//--------------------------------------------------------------------Setters

// 直接给对象的引用类型实例变量赋值
func (self *Object) SetRefVar(name string, descriptor string, ref *Object) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetRef(field.slotId, ref)
}

func (self *Object) SetExtra(extra interface{}) {
	self.extra = extra
}

//--------------------------------------------------------------------功能类方法

func (self *Object) Clone() *Object {
	return &Object{
		class: self.class,
		data:  self.cloneData(),
	}
}

func (self *Object) cloneData() interface{} {
	switch self.data.(type) {
	case []int8:
		elements := self.data.([]int8)
		elements2 := make([]int8, len(elements))
		copy(elements2, elements)
		return elements2
	case []int16:
		elements := self.data.([]int16)
		elements2 := make([]int16, len(elements))
		copy(elements2, elements)
		return elements2
	case []uint16:
		elements := self.data.([]uint16)
		elements2 := make([]uint16, len(elements))
		copy(elements2, elements)
		return elements2
	case []int32:
		elements := self.data.([]int32)
		elements2 := make([]int32, len(elements))
		copy(elements2, elements)
		return elements2
	case []int64:
		elements := self.data.([]int64)
		elements2 := make([]int64, len(elements))
		copy(elements2, elements)
		return elements2
	case []float32:
		elements := self.data.([]float32)
		elements2 := make([]float32, len(elements))
		copy(elements2, elements)
		return elements2
	case []float64:
		elements := self.data.([]float64)
		elements2 := make([]float64, len(elements))
		copy(elements2, elements)
		return elements2
	case []*Object:
		elements := self.data.([]*Object)
		elements2 := make([]*Object, len(elements))
		copy(elements2, elements)
		return elements2
	default: // []Slot
		slots := self.data.(Slots)
		slots2 := newSlots(uint(len(slots)))
		copy(slots2, slots)
		return slots2
	}
}

func (self *Object) GetData() interface{} {
	return self.data
}

func (self *Object) SetIntVar(name string, descriptor string, value int32) {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	slots.SetInt(field.slotId, value)
}

func (self *Object) GetIntVar(name, descriptor string) int32 {
	field := self.class.getField(name, descriptor, false)
	slots := self.data.(Slots)
	return slots.GetInt(field.slotId)
}

func ArrayCopy(src, dst *Object, srcPos, dstPos, length int32) {
	switch src.data.(type) {
	case []int8:
		_src := src.data.([]int8)[srcPos : srcPos+length]
		_dst := dst.data.([]int8)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int16:
		_src := src.data.([]int16)[srcPos : srcPos+length]
		_dst := dst.data.([]int16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int32:
		_src := src.data.([]int32)[srcPos : srcPos+length]
		_dst := dst.data.([]int32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []int64:
		_src := src.data.([]int64)[srcPos : srcPos+length]
		_dst := dst.data.([]int64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []uint16:
		_src := src.data.([]uint16)[srcPos : srcPos+length]
		_dst := dst.data.([]uint16)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float32:
		_src := src.data.([]float32)[srcPos : srcPos+length]
		_dst := dst.data.([]float32)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []float64:
		_src := src.data.([]float64)[srcPos : srcPos+length]
		_dst := dst.data.([]float64)[dstPos : dstPos+length]
		copy(_dst, _src)
	case []*Object:
		_src := src.data.([]*Object)[srcPos : srcPos+length]
		_dst := dst.data.([]*Object)[dstPos : dstPos+length]
		copy(_dst, _src)
	default:
		panic("Not array!")
	}
}

