package state

import (
	"fmt"
	"github.com/tianbingJ/lua-go/lua-dump/vm"
	"tianbingj.github.com/binchunk"
)

//return 0表示成功，非0表示失败
//mode = 'b', 二进制； =='t'，脚本； == 'bt'，二者皆可
func (self *luaState) Load(chunk []byte, chunkName, mode string) int {
	proto := binchunk.Undump(chunk)
	c := newLuaClosure(proto)
	self.stack.push(c)
	return 0
}

//调用函数之前，要把closure和参数准备好
//栈里的内容是函数和参数
//函数f， 参数a、b、c；栈结构是：[f, a, b, c]
//Call执行结束后，函数和参数会从栈里弹出，返回值放入栈中
//nArgs参数的数量，能推断出函数在栈里的位置
//nResults保留的函数调用结果数量，如果是-1则表示返回结果全部保存在栈里
func (self *luaState) Call(nArgs, nResults int) {
	val := self.stack.get(-(nArgs + 1))
	if c, ok := val.(*closure); ok {
		fmt.Printf("call %s<%d,%d>\n", c.proto.Source, c.proto.LineDefined, c.proto.LastLineDefined)
		self.callLuaClosure(nArgs, nResults, c)
	} else {
		panic("not function!")
	}
}

func (self *luaState) callLuaClosure(nArgs, nResults int, c *closure) {
	nRegs := int(c.proto.MaxStackSize)
	//固定参数数量
	nParams := int(c.proto.NumParams)
	isVararg := c.proto.IsVararg == 1

	newStack := newLuaStack(nRegs + 20)
	newStack.closure = c

	//从栈中弹出函数和参数
	funcAndArgs := self.stack.popN(nArgs + 1)
	newStack.pushN(funcAndArgs[1:], nParams)
	newStack.top = nRegs
	if nArgs > nParams && isVararg {
		newStack.varargs = funcAndArgs[nParams+1:]
	}

	self.pushLuaStack(newStack)
	self.runLuaClosure()
	self.popLuaStack()

	if nResults != 0 {
		//从被调用帧弹出，放入到调用帧中；为什么弹出newStack.top - nRegs个？
		results := newStack.popN(newStack.top - nRegs)
		self.stack.check(len(results))
		self.stack.pushN(results, nResults)
	}
}

func (self *luaState) runLuaClosure() {
	for {
		inst := vm.Instruction(self.Fetch())
		inst.Execute(self)
		if inst.Opcode() == vm.OP_RETURN {
			break
		}
	}
}
