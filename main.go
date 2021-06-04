package main

import "syscall/js"

func add(this js.Value, args []js.Value) interface{} {
	x := args[0].Int()
	y := args[1].Int()
	return js.ValueOf(x + y)
}

func main() {
	done := make(chan struct{}, 0)

	addFunc := js.FuncOf(add)
	js.Global().Set("addFunc", addFunc)
	defer addFunc.Release()

	<-done
}
