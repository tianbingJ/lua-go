package vm

import "github.com/tianbingJ/lua-go/lua-dump/api"

const LFIELDS_PER_FLUSH = 50

//创建新的表，放入在A寄存器中，size = B, C
func newTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.CreateTable(Int2fb(b), Int2fb(c))
	vm.Replace(a)
}

//R(A) := R(B) [RK(C)]
//从B指定的表里R(B)取出数据，写入目标寄存器中R(A);Key可能在寄存器中，也可能在
func getTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	vm.GetRK(c)
	vm.GetTable(b)
	vm.Replace(a)
}

//R(A)[RK(B)) := RK(C)
func setTable(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	vm.GetRK(b)
	vm.GetRK(c)
	vm.SetTable(a)
}

//R(A)[(C - 1) *FPF + i] := R(A+i)   1 <=i <= B
//设置table表的list部分, Table由A指定；数量由B指定，数组的起始索引由C指定。
//C寄存器最多只有9位，表示的数量有限，不够表示数组的索引。数组的C保存的实际是批次数量。
//默认的批次大小是50,C操作数的最大索引是：50 * 512 = 25600
//如果C不够用，则使用扩展执行EXTRAARG，读取Ax里的信息
func setList(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	if c > 0 {
		c -= 1
	} else {
		c = Instruction(vm.Fetch()).Ax()
	}

	idx := int64(c * LFIELDS_PER_FLUSH)
	for j := 1; j <= b; j ++ {
		idx ++
		vm.PushValue(a + j)
		vm.SetI(a, idx)
	}
}
