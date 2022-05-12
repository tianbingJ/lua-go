package api

type LuaVM interface {
	LuaState
	PC() int
	AddPC(n int)      //修改PC，用于跳转
	Fetch() uint32    //取出当前指令；将PC指向下一条指令
	GetConst(idx int) //将指定常量推入栈顶
	GetPK(rk int)     //将指定常量或栈值推入栈顶
}
