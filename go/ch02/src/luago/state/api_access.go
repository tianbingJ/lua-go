package state

import (
	"fmt"
	"github.com/tianbingJ/lua-go/lua-dump/api"
)

//access方法用于从栈中获取信息

func (self *luaState) TypeName(tp api.LuaType) string {
	switch tp {
	case api.LUA_TNONE:
		return "no value"
	case api.LUA_TNIL:
		return "nil"
	case api.LUA_TSTRING:
		return "string"
	case api.LUA_TBOOLEAN:
		return "boolean"
	case api.LUA_TNUMBER:
		return "number"
	case api.LUA_TTABLE:
		return "table"
	case api.LUA_TTHREAD:
		return "thread"
	case api.LUA_TFUNCTION:
		return "function"
	default:
		return "userdata"
	}
}

func (self *luaState) Type(idx int) api.LuaType {
	if self.stack.isValid(idx) {
		val := self.stack.get(idx)
		return typeOf(val)
	}
	return api.LUA_TNONE
}

func (self *luaState) IsNone(idx int) bool {
	return self.Type(idx) == api.LUA_TNONE
}

func (self *luaState) IsNil(idx int) bool {
	return self.Type(idx) == api.LUA_TNIL
}

func (self *luaState) IsNoneOrNil(idx int) bool {
	return self.IsNone(idx) || self.IsNil(idx)
}

func (self *luaState) IsBoolean(idx int) bool {
	return self.Type(idx) == api.LUA_TBOOLEAN
}

//这里其实没有看懂，为什么要判断是否是number?
func (self *luaState) IsString(idx int) bool {
	t := self.Type(idx)
	return t == api.LUA_TSTRING || t == api.LUA_TNUMBER
}

func (self *luaState) IsNumber(idx int) bool {
	_, ok := self.ToNumberX(idx)
	return ok
}

func (self *luaState) IsInteger(idx int) bool {
	val := self.stack.get(idx)
	_, ok := val.(int64)
	return ok
}

func (self *luaState) ToBoolean(idx int) bool {
	val := self.stack.get(idx)
	return convertToBoolean(val)
}

//lua中只有nil和false表示假
func convertToBoolean(val luaValue) bool {
	switch x := val.(type) {
	case nil:
		return false
	case bool:
		return x
	default:
		return true
	}
}

func (self *luaState) ToNumber(idx int) float64 {
	n, _ := self.ToNumberX(idx)
	return n
}

func (self *luaState) ToNumberX(idx int) (float64, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case float64:
		return float64(x), true
	case int64:
		return float64(x), true
	default:
		return 0, false
	}
}

func (self *luaState) ToInteger(idx int) int64 {
	v, _ := self.ToIntegerX(idx)
	return v
}

func (self *luaState) ToIntegerX(idx int) (int64, bool) {
	val := self.stack.get(idx)
	i, ok := val.(int64)
	return i, ok
}

func (self *luaState) ToStringX(idx int) (string, bool) {
	val := self.stack.get(idx)
	switch x := val.(type) {
	case string:
		return x, true
	case int64, float64:
		s := fmt.Sprintf("%v", x)
		self.stack.set(idx, s)
		return s, true
	default:
		return "", false
	}
}

func (self *luaState) ToString(idx int) string {
	s, _ := self.ToStringX(idx)
	return s
}