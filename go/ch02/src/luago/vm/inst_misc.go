package vm

import "github.com/tianbingJ/lua-go/lua-dump/api"

//R(A) := R(B)
//寄存器的值+1对应到栈的索引
func move(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Copy(b, a)
}

func jmp(i Instruction, vm api.LuaVM) {
	a, sBx := i.AsBx()
	vm.AddPC(sBx)
	if a != 0 {
		panic("todo!")
	}
}