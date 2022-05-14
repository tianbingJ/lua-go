package vm

import "github.com/tianbingJ/lua-go/lua-dump/api"

//forprep
//R(A) -= R(A+2); pc += sBx
//A是循环控制变量， A + 2是步长
func forPrep(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	//R(A) -= R(A+2)
	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPSUB)
	vm.Replace(a)
	vm.AddPC(sBx)
}

//forloop
//R(A) += R(A+2)
//if R(A) <?= R(A + 1) { pc += sBx; R(A + 3) = R(A)}
//	当步长是> 0时， <？=表示 <=; 当步长 < 0时，<?=表示 >=

func forLoop(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	a += 1

	//R(A) += R(A+2)
	vm.PushValue(a)
	vm.PushValue(a + 2)
	vm.Arith(api.LUA_OPADD)
	vm.Replace(a)

	isPositiveStep := vm.ToNumber(a+2) >= 0
	if isPositiveStep && vm.Compare(a, a+1, api.LUA_OPLE) || !isPositiveStep && vm.Compare(a + 1, a, api.LUA_OPLE) {
		vm.AddPC(sBx)
		vm.Copy(a, a + 3)
	}
}