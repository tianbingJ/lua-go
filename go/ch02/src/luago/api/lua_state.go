package api

type LuaType = int
type ArithOp = int
type CompareOp = int

type LuaState interface {
	GetTop() int
	AbsIndex(idx int) int
	CheckStack(n int) bool
	Pop(n int)
	Copy(fromIdx, toIdx int)
	PushValue(idx int)
	//用栈顶的值替换掉idx位置的值
	Replace(idx int)
	Insert(idx int)
	Remove(idx int)
	Rotate(idx int, n int)
	SetTop(idx int)

	TypeName(tp LuaType) string
	Type(idx int) LuaType
	IsNone(idx int) bool
	IsNil(idx int) bool
	IsNoneOrNil(idx int) bool
	IsBoolean(idx int) bool
	IsInteger(idx int) bool
	IsNumber(idx int) bool
	IsString(idx int) bool
	ToBoolean(idx int) bool
	ToInteger(idx int) int64
	ToIntegerX(idx int) (int64, bool)
	ToNumber(idx int) float64
	ToNumberX(idx int) (float64, bool)
	ToString(idx int) string
	ToStringX(idx int) (string, bool)

	PushNil()
	PushBoolean(b bool)
	PushInteger(n int64)
	PushNumber(n float64)
	PushString(s string)

	Arith(op ArithOp)
	Compare(idx1, idx2 int, op CompareOp) bool
	Len(idx int)
	Concat(n int)

	//表相关api
	NewTable()
	CreateTable(nArr, nReg int)
	//获取idx位置的table，并返回table[stack[top]]
	GetTable(idx int) LuaType
	GetField(idx int, k string) LuaType
	GetI(idx int, i int64) LuaType
	SetTable(idx int)
	SetField(idx int, k string)
	SetI(idx int, n int64)
}
