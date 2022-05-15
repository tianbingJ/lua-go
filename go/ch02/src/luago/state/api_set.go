package state

//key和value在栈中
func (self *luaState) SetTable(idx int) {
	t := self.stack.get(idx)
	v := self.stack.pop()
	k := self.stack.pop()
	self.setTable(t, k, v)
}

func (self *luaState) setTable(table, key, value luaValue) {
	if tbl, ok := table.(*luaTable); ok {
		tbl.put(key, value)
		return
	}
	panic("not table")
}

//键不是从栈中弹出，而是由参数传入
func (self *luaState) setField(idx int, k string) {
	t := self.stack.get(idx)
	value := self.stack.pop()
	self.setTable(t, k, value)
}

func (self *luaState) setFielI(idx int, i int64) {
	t := self.stack.get(idx)
	value := self.stack.pop()
	self.setTable(t, i, value)
}
