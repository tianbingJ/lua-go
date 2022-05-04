package binchunk

const (
	LUA_SIGNATURE        = "\x1bLua"
	LUAC_VERSION         = 0x53
	LUAC_FORMAT          = 0
	LUAC_DATA            = "\x19\x93\r\n\x1a\n"
	CINT_SIZE            = 4
	CSIZET_SIZE          = 8
	LUA_INSTRUCTION_SIZE = 4
	LUA_INTEGER_SIZE     = 8
	LUA_NUMBER_SIZE      = 8
	LUAC_INT             = 0x5678
	LUAC_NUM             = 370.5
)

type binaryChunk struct {
	header
	sizeUpvalues byte
	mainFunc     *Prototype
}

type header struct {
	signature       [4]byte //魔数 ESC L u a,0X1B4C7561
	version         byte    //大版本号 5，小版本号 3;5.3对应的版本是83，16进制53
	format          byte    //固定值 0
	luacData        [6]byte //前两个是 0x1993，lua发型的年份; 后四个是回车(0x0D)、换行符(0x0A)、替换符(0x1A)、换行符(0x0A)
	cintSize        byte    //cint,列表长度占用的字节数:4
	sizetSize       byte    //sizet，字符串长度占用的字节数:8
	instructionSize byte    //指令长度占用字节数 4
	luaIntegerSize  byte    //整数占用字节数 8
	luaNumberSize   byte    //浮点数占用字节数 8
	luacInt         int64   //整数0x5678，由于int占用8个字节，所以结果是78 56 00 00 00 00 00 00，检测大小端对应方式
	luacNum         float64 //存放浮点数 370.5，占用8个字节；检测浮点数格式与本地是否匹配；IEEE 754浮点数格式
}

//函数原型
type Prototype struct {
	Source          string //文件名 '@'后是真正的文件名，表示从源文件而来; '=stdin'表示从标准输入编译而来
	LineDefined     uint32 //记录原型在源文件中的起止行号；普通的函数，起止行号都大于0；主函数，起止行号都是0
	LastLineDefined uint32
	NumParams       byte          //固定参数个数
	IsVararg        byte          //是否有变长参数0表示否，1表示是；主函数有，值是1
	MaxStackSize    byte          //寄存器的数量
	Code            []uint32      //指令表，每条指令占用4个字节
	Constants       []interface{} //常量表，每个常量都以1字节tag打头，用来标识后面存储的是哪种类型的常量
	Upvalues        []Upvalue     //每个占用2个字节
	Protos          []*Prototype  //子函数原型列表
	LineInfo        []uint32      //与指令表里的数据一一对应
	LocVars         []LocVar      //局部变量表
	UpvalueNames    []string	 //和upvalues列表一一对应
}

type Upvalue struct {
	Instack byte
	Idx     byte
}

type LocVar struct {
	VarName string
	StartPc uint32
	EndPC   uint32
}

//常量池打头的tag
const (
	TAG_NIL       = 0x00
	TAG_BOOLEAN   = 0x01
	TAG_NUMBER    = 0x03
	TAG_INTEGER   = 0x13
	TAG_SHORT_STR = 0x04
	TAG_LONG_STR  = 0x14
)

//解析二进制chunk
func Undump(data []byte) *Prototype {
	reader := &reader{data}
	reader.checkHeader()
	reader.readByte()
	return reader.readProto("")
}