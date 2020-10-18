package rtda

import (
	"jvm_go_code/invoke_return/rtda/heap"
	"math"
	"strconv"
)

type OperandStack struct {
	size	uint    // 容量
	slots	[]Slot // 插槽
}

func newOperandStack(max uint) *OperandStack {
	if max > 0 {
		return &OperandStack {
			slots: make([]Slot, max),
		}
	}
	return nil
}

func (self *OperandStack) PushInt(val int32) {
	self.slots[self.size].num = val
	self.size++

}

func (self *OperandStack) PopInt() int32 {
	self.size--
	result := self.slots[self.size]
	return result.num
}

func (self *OperandStack) PushFloat(val float32)  {
	bits := math.Float32bits(val)
	self.slots[self.size].num = int32(bits)
	self.size++
}

func (self *OperandStack) PopFloat() float32 {
	self.size--
	bits := uint32(self.slots[self.size].num)
	return float32(bits)
}

func (self *OperandStack) PushLong(val int64) {
	self.slots[self.size].num = int32(val)
	self.slots[self.size + 1].num = int32(val >> 32)
	self.size += 2
}

func (self *OperandStack) PopLong() int64 {
	self.size -= 2
	low := uint32(self.slots[self.size].num)
	high := uint32(self.slots[self.size + 1].num)
	return int64(high) << 32 | int64(low)
}

func (self *OperandStack) PushDouble(val float64) {
	bits := math.Float64bits(val)
	self.PushLong(int64(bits))
}

func (self *OperandStack) PopDouble() float64 {
	bits := uint64(self.PopLong())
	return math.Float64frombits(bits)
}

func (self *OperandStack) PushRef(val *heap.Object) {
	self.slots[self.size].ref = val
	self.size++
}

func (self *OperandStack) PopRef() *heap.Object {
	self.size--
	result := self.slots[self.size].ref
	self.slots[self.size].ref = nil
	return result
}


func (self *OperandStack) PushSlot(slot Slot)  {
	self.slots[self.size] = slot
	self.size++
}

func (self *OperandStack) PopSlot() Slot {
	self.size--
	return self.slots[self.size]
}

func (self OperandStack) String() string {
	var result string = "size: " + strconv.Itoa(int(self.size)) + "; "
	for index, item := range self.slots {
		result += item.String()
		if index != len(self.slots) - 1 {
			result += ", "
		}
	}
	return result
}