package heap

import (
	"math"
)

type Slot struct {
	num int32
	ref *Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	if slotCount > 0 {
		return make([]Slot, slotCount)
	}
	return nil
}

func (self Slots) SetInt(index uint, value int32) {
	self[index].num = value
}

func (self Slots) GetInt(index uint) int32 {
	return self[index].num
}

func (self Slots) SetFloat(index uint, value float32) {
	bits := math.Float32bits(value)
	self[index].num = int32(bits)
}

func (self Slots) GetFloat(index uint) float32 {
	bits := uint32(self[index].num)
	return math.Float32frombits(bits)
}

func (self Slots) SetLong(index uint, value int64) {
	self[index].num = int32(value)
	self[index + 1].num = int32(value >> 32)
}

func (self Slots) GetLong(index uint) int64 {
	low := uint32(self[index].num)
	high := uint32(self[index + 1].num)
	return int64(high) << 32 | int64(low)
}

func (self Slots) SetDouble(index uint, value float64) {
	bits := math.Float64bits(value)
	self.SetLong(index, int64(bits))
}

func (self Slots) GetDouble(index uint) float64 {
	bits := uint64(self.GetLong(index))
	return math.Float64frombits(bits)
}

func (self Slots) SetRef(index uint, value *Object) {
	self[index].ref = value
}

func (self Slots) GetRef(index uint) *Object {
	return self[index].ref
}