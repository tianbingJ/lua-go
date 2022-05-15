package state

import "github.com/tianbingJ/lua-go/lua-dump/api"

func (self *luaState) CreateTable(nArr, nReg int) {
	t := newLuaTable(nArr, nReg)
	self.stack.push(t)
}

func (self *luaState) NewTable() {
	self.CreateTable(0, 0)
}

func (self *luaState) GetTable(idx int) api.LuaType {
	t := self.stack.get(idx)
	k := self.stack.pop()
	return self.getTable(t, k)
}

func (self *luaState) getTable(t, k luaValue) api.LuaType {
	if tbl, ok := t.(*luaTable); ok {
		v := tbl.get(k)
		self.stack.push(v)
		return typeOf(v)
	}
	panic("not a table")
}

//与GetTable不同的是，key是由参数指定，而不是在栈里
func (self *luaState) GetField(idx int, k string) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, k)
}

func (self *luaState) GetI(idx int, i int64) api.LuaType {
	t := self.stack.get(idx)
	return self.getTable(t, i)
}


