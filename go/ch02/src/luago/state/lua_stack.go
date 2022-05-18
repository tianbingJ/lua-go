package state

//valid indices:指向可修改的栈位置，包括[1, 到栈顶n] (1 <= abs(idx) <= top)和 pseudo-indices(可以被外部语言访问，但不在栈里)
//acceptable indices:
//1.可以是valid indices,也包括给栈分配了空间但是在栈顶之外的位置； 0不是一个acceptable index
//2.index < 0 && abs(index) <= top || (index > 0 && index <= stackspace)

//lua的堆外的栈索引是从1开始, [1,n]
//但是这个top指向的是栈顶元素的下一个位置?
type luaStack struct {
	slots   []luaValue //存放值
	top     int
	prev    *luaStack
	closure *closure
	varargs []luaValue
	pc      int
}

func newLuaStack(size int) *luaStack {
	return &luaStack{
		slots: make([]luaValue, size),
		top:   0,
	}
}

//检查是否能够有n个容量，如果不足，则进行扩容
func (self *luaStack) check(size int) {
	free := self.top - len(self.slots)
	for i := free; i < size; i++ {
		self.slots = append(self.slots, nil)
	}
}

func (self *luaStack) push(val luaValue) {
	if self.top == len(self.slots) {
		panic("stack overflow!")
	}
	self.slots[self.top] = val
	self.top++
}

func (self *luaStack) pop() luaValue {
	if self.top < 1 {
		panic("stack underflow!")
	}
	self.top--
	val := self.slots[self.top]
	self.slots[self.top] = nil
	return val
}

//把索引转为绝对索引, 不检查有效性
//A positive index represents an absolute stack position (starting at 1);
func (self *luaStack) absIndex(idx int) int {
	if idx >= 0 {
		return idx
	}
	//-1是栈顶
	return idx + self.top + 1
}

func (self *luaStack) isValid(idx int) bool {
	absIdx := self.absIndex(idx)
	return absIdx > 0 && absIdx <= self.top
}

func (self *luaStack) get(idx int) luaValue {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		return self.slots[absIdx-1]
	}
	return nil
}

func (self *luaStack) set(idx int, val luaValue) {
	absIdx := self.absIndex(idx)
	if absIdx > 0 && absIdx <= self.top {
		self.slots[absIdx-1] = val
		return
	}
	panic("invalid index!")
}

func (self *luaStack) reverse(from, to int) {
	slots := self.slots
	for from < to {
		slots[from], slots[to] = slots[to], slots[from]
		from++
		to--
	}
}
