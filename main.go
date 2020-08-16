package main

import lua "github.com/yuin/gopher-lua"

func main() {
	L := lua.NewState()
	defer L.Close()
	if err := L.DoFile("foo.lua"); err != nil {
		panic(err)
	}
}
