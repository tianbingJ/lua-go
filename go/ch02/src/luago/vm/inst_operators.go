package vm

import "github.com/tianbingJ/lua-go/lua-dump/api"

//R(A) = RK(B) op RK(C)
func _binaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, c := i.ABC()
	vm.GetRK(b)
	vm.GetRK(c)
	vm.Arith(op)
	vm.Replace(a)
}

// R(a) := op R(B)
func _unaryArith(i Instruction, vm api.LuaVM, op api.ArithOp) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushValue(b)
	vm.Arith(op)
	vm.Replace(a)
}

func add(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPADD)
}

func sub(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPSUB)
}

func mul(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPMUL)
}

func mod(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPMOD)
}

func pow(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPPOW)
}

func div(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPDIV)
}

func idiv(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPIDIV)
}

func band(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPBAND)
}

func bor(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPBOR)
}

func bxor(i Instruction, vm api.LuaVM) {
	_binaryArith(i, vm, api.LUA_OPBXOR)
}

func shl(i Instruction, vm api.LuaVM)  {
	_binaryArith(i, vm, api.LUA_OPSHL)
}

func shr(i Instruction, vm api.LuaVM) {
	 _binaryArith(i, vm, api.LUA_OPSHR)
}

func unm(i Instruction, vm api.LuaVM) {
	_unaryArith(i, vm, api.LUA_OPUNM)
}

func bnot(i Instruction, vm api.LuaVM) {
	_unaryArith(i, vm, api.LUA_OPBNOT)
}


//R(A) := len(R(B))
func len(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.Len(b)
	vm.Replace(a)
}

//R(A) := R(B) ..R(B+1).. ... .. R(C)
func concat(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	c += 1
	n := c - b + 1
	vm.CheckStack(n)
	for i := b; i <= c; i ++ {
		vm.PushValue(i)
	}
	vm.Concat(n)
	vm.Replace(a)
}

//if (RK(B) op RK(C) ~= A) then pc ++
func _compare(i Instruction, vm api.LuaVM, op api.CompareOp) {
	a, b, c := i.ABC()
	vm.GetRK(b)
	vm.GetRK(c)
	r := a != 0
	if vm.Compare(-2, -1, op) != r {
		vm.AddPC(1)
	}
	vm.Pop(2)
}

func eq(i Instruction, vm api.LuaVM) {
	_compare(i, vm, api.LUA_OPEQ)
}

func lt(i Instruction, vm api.LuaVM) {
	_compare(i, vm, api.LUA_OPLT)
}

func le(i Instruction, vm api.LuaVM) {
	_compare(i, vm, api.LUA_OPLE)
}

//R(A) := not R(B)
func not(i Instruction, vm api.LuaVM) {
	a, b, _ := i.ABC()
	a += 1
	b += 1
	vm.PushBoolean(vm.ToBoolean(b))
	vm.Replace(a)
}

//if (R(B) == C) then R(A) := R(B) else pc ++
func testSet(i Instruction, vm api.LuaVM) {
	a, b, c := i.ABC()
	a += 1
	b += 1
	cb := c !=0
	if vm.ToBoolean(b) == cb {
		vm.Copy(b, a)
	} else {
		vm.AddPC(1)
	}
}

//if R(A) != C then pc ++
func test(i Instruction, vm api.LuaVM) {
	a, _, c := i.ABC()
	a += 1
	cb := c != 0
	if vm.ToBoolean(a) != cb {
		vm.AddPC(1)
	}
}