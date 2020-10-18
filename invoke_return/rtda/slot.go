package rtda

import (
	"jvm_go_code/invoke_return/rtda/heap"
	"strconv"
)

type Slot struct {
	num		int32        // 基本数据类型
	ref		*heap.Object // 引用数据类型
}

func (self *Slot) String() string {
	return "[num: " + strconv.Itoa(int(self.num)) + ", ref: " + self.ref.String() + "]"
}
