package state

import (
	"github.com/tianbingJ/lua-go/lua-dump/api"
	"github.com/tianbingJ/lua-go/lua-dump/number"
	"math"
)

var (
	iadd = func(a, b int64) int64 {
		return a + b
	}
	fadd = func(a, b float64) float64 {
		return a + b
	}
	isub = func(a, b int64) int64 {
		return a - b
	}
	fsub = func(a, b float64) float64 {
		return a - b
	}
	imul = func(a, b int64) int64 {
		return a * b
	}
	fmul = func(a, b float64) float64 {
		return a * b
	}
	imod = number.IMod
	fmod = number.FMod
	pow  = math.Pow
	div  = func(a, b float64) float64 {
		return a / b
	}
	iidiv = number.IFloorDiv
	ffdiv = number.FFloorDiv
	band  = func(a, b int64) int64 {
		return a & b
	}
	bor = func(a, b int64) int64 {
		return a | b
	}
	bxor = func(a, b int64) int64 {
		return a ^ b
	}
	shl  = number.ShiftLeft
	shr  = number.ShiftRight
	iunm = func(a, _ int64) int64 {
		return -a
	}
	funm = func(a, _ float64) float64 {
		return -a
	}
	bnot = func(a, _ int64) int64 {
		return ^a
	}
)

type operator struct {
	integerFunc func(int64, int64) int64
	floatFunc   func(float64, float64) float64
}

//nil表示不支持对应类型的运算；
var operators = []operator{
	{iadd, fadd},
	{isub, fsub},
	{imul, fmul},
	{imod, fmod},
	{nil, pow},
	{nil, div},
	{iidiv, ffdiv},
	{band, nil},
	{bor, nil},
	{bxor, nil},
	{shl, nil},
	{shr, nil},
	{iunm, funm},
	{bnot, nil},
}

func (self *luaState) Arith(op api.ArithOp) {
	var a, b luaValue
	b = self.stack.pop()
	if op != api.LUA_OPUNM && op != api.LUA_OPBNOT {
		a = self.stack.pop()
	} else {
		a = b
	}
	operator := operators[op]
	if result := _arith(a, b, operator); result != nil {
		self.stack.push(result)
	} else {
		panic("arithmetic error!")
	}
}

func _arith(a, b luaValue, op operator) luaValue {
	if op.floatFunc == nil {
		if x, ok := convertToInteger(a); ok {
			if y, ok := convertToInteger(b); ok {
				return op.integerFunc(x, y)
			}
		}
	} else {
		if op.integerFunc != nil {
			if x, ok := a.(int64); ok {
				if y, ok := b.(int64); ok {
					return op.integerFunc(x, y)
				}
			}
		}
		if x, ok := convertToFloat(a); ok {
			if y, ok := convertToFloat(b); ok {
				return op.floatFunc(x, y)
			}
		}
	}
	return nil
}

