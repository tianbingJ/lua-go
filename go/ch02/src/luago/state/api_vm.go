package state

func (self *luaState) PC() int {
	return self.stack.pc
}

func (self *luaState) AddPC(n int) {
	self.stack.pc += n
}

func (self *luaState) Fetch() uint32 {
	i := self.stack.closure.proto.Code[self.stack.pc]
	self.stack.pc ++
	return i
}

func (self *luaState) GetConst(idx int) {
	c := self.stack.closure.proto.Constants[idx]
	self.stack.push(c)
}

//根据情况调用GetConst()方法把某个常量推入栈顶，或者调用pushValue把某个索引的值推入栈顶
func (self *luaState) GetRK(rk int) {
	if rk > 0xFF { //const
		self.GetConst(rk & 0xFf)
	} else {
		self.PushValue(rk + 1)
	}
}