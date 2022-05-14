package vm

import "github.com/tianbingJ/lua-go/lua-dump/api"


//R(A), R(A+1),..., R(A + B) := nil
//起始寄存器由A决定，数量由寄存器B决定
func loadNil(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1

	vm.PushNil()
	for i := a; i <= a+b; i++ {
		vm.Copy(-1, i)
	}
	vm.Pop(1)
}

//R(A) := (Bool)B; if (C) pc ++
func loadBool(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.PushBoolean(b != 0)
	vm.Replace(a)
	if c != 0 {
		vm.AddPC(1)
	}
}

//将常量池中的某个常量加载到指定寄存器中,寄存器索引由操作数A表示，常量表由操作数Bx表示
//R(A) := Kst(N)

// >luac -l -
// local a,b = 1, "foo"
//0+ params, 2 slots, 1 upvalue, 2 locals, 2 constants, 0 functions
//	1	[1]	LOADK    	0 -1	; 1
//	2	[1]	LOADK    	1 -2	; "foo"
//	3	[1]	RETURN   	0 1
// 为什么luac中常量池的下标是 负数？
// 从常量池展示的是负数，与从寄存器里区分。
func loadK(i Instruction, vm api.LuaVM) {
	a, bx := i.ABx()
	a += 1
	vm.GetConst(bx)
	vm.Replace(a)
}

//loadK的Bx 18位，支持的常量池最大262143个; loadKx支持更大的常量池。
//loadKx需要与EXTRAARG指令配合使用，用后者的Ax操作数指定常量池索引。
func loadKx(i Instruction, vm api.LuaVM) {
	a, _ := i.ABx()
	a += 1
	ax := Instruction(vm.Fetch()).Ax()

	vm.GetConst(ax)
	vm.Replace(a)
}