package state

import "tianbingj.github.com/binchunk"

type luaState struct {
	stack *luaStack
	proto *binchunk.Prototype
	pc int
}

func New(stackSize int, proto *binchunk.Prototype) *luaState {
	return &luaState{
		stack: newLuaStack(stackSize),
		pc: 0,
		proto: proto,
	}
}

func (self *luaState) GetTop() int {
	return self.stack.top
}

func (self *luaState) AbsIndex(idx int) int {
	return self.stack.absIndex(idx)
}

func (self *luaState) CheckStack(n int) bool {
	self.stack.check(n)
	return true
}

func (self *luaState) Pop(n int) {
	for i := 0; i < n; i++ {
		self.stack.pop()
	}
}

func (self *luaState) Copy(fromIdx, toIdx int) {
	val := self.stack.get(fromIdx)
	self.stack.set(toIdx, val)
}

//指定位置的值再次压入栈
func (self *luaState) PushValue(idx int) {
	val := self.stack.get(idx)
	self.stack.push(val)
}

//栈顶值弹出，然后写入指定位置
func (self *luaState) Replace(idx int) {
	val := self.stack.pop()
	self.stack.set(idx, val)
}

//将[idx, top]索引区间内的值超栈顶旋转n个位置
//n > 0 向上旋转
// n < 0,向下旋转
func (self *luaState) Rotate(idx, n int) {
	//t，栈顶部下标
	//p, 旋转起始位置下标
	//m, 选择位置的下标
	t := self.stack.top - 1
	p := self.stack.absIndex(idx) - 1
	var m int
	if n >= 0 {
		m = t - n
	} else {
		m = p - n - 1
	}
	self.stack.reverse(p, m)
	self.stack.reverse(m+1, t)
	self.stack.reverse(p, t)
}

//栈顶弹出,插入在idx处
func (self *luaState) Insert(idx int) {
	self.Rotate(idx, 1)
}

func (self *luaState) Remove(idx int) {
	self.Rotate(idx, -1)
	self.Pop(1)
}

//将idx指向的栈顶设置为top
//如果idx小于栈顶，则相当于执行多次pop
//如果idx > 栈顶，则相当于push多个nil
func (self *luaState) SetTop(idx int) {
	newTop := self.AbsIndex(idx)
	if newTop < 0 {
		panic("stack underflow!")
	}
	n := self.stack.top - newTop
	if n > 0 {
		self.Pop(n)
	} else if n < 0 {
		for i := 0; i < - n; i ++ {
			self.stack.push(nil)
		}
	}
}
