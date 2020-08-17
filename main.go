package main

import (
	"fmt"

	lua "github.com/yuin/gopher-lua"
)

func main() {
	// Set up Lua State and load file
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("foo.lua"); err != nil {
		panic(err)
	}

	// Call concat
	if err := L.CallByParam(lua.P{
		Fn:      L.GetGlobal("concat"), // function name
		NRet:    1,                     // number of returns
		Protect: true,                  // return err or panic
	}, lua.LString("Go"), lua.LString("Lua")); err != nil {
		panic(err)
	}

	// Get the returned value from the stack and cast it to a Lua string
	if str, ok := L.Get(-1).(lua.LString); ok {
		fmt.Println(str)
	}

	// Pop return value off stack
	L.Pop(1)
}
