package heap

import "log"

type Object struct {
	class			*Class
	fields			Slots
}

func newObject(class *Class) *Object {
	log.Println(class.instanceSlotCount)
	return &Object {
		class: class,
		fields: newSlots(class.instanceSlotCount),
	}
}

func (self *Object) Class() *Class {
	return self.class
}

func (self *Object) Fields() Slots {
	return self.fields
}

func (self *Object) IsInstanceOf(class *Class) bool {
	return class.isAssignableFrom(self.class)
}

func (self *Object) String() string {
	return "nil"
}