package govm

import (
	"testing"
	"./opcode"
	"./types"
)

func TestGovm(t *testing.T) {
	// See hello.gvm for a more readable version
	code := []byte{
		byte(opcode.Func),
		0x00, 0x00, 0x00, 0x00, // 0 args
		0x00, 0x00, 0x00, 0x00, // 0 return values
		0x00, 0x00, 0x00, 0x27, // 39 bytes
		byte(opcode.Push), byte(types.String),
		0x00, 0x00, 0x00, 0x0d, // 13 bytes
	}
	code = append(code, []byte("Hello, world!")...)
	code = append(code,
		byte(opcode.Get),
		0x00, 0x00, 0x00, 0x0e, // 14 bytes
	)
	code = append(code, []byte("Println:string")...)
	code = append(code,
		byte(opcode.Call),
		byte(opcode.Set),
		0x00, 0x00, 0x00, 0x05, // 5 bytes
	)
	code = append(code, []byte("Main:")...)

	v := New()
	v.Load(code)

	v.Get(types.Symbol("Main:"))
	v.Call()
}
