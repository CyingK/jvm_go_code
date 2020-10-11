package rtda

import "strconv"

type Slot struct {
	num		int32		// 基本数据类型
	ref		*Object	// 引用数据类型
}

func (self *Slot) String() string {
	return "[num: " + strconv.Itoa(int(self.num)) + ", ref: " + self.ref.String() + "]"
}
