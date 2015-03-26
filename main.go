package main

import (
    "fmt"
    "github.com/nsf/gothic"
    )

func main() {
	ir := gothic.NewInterpreter(buildAndInitApp)
	<-ir.Done
}

func buildAndInitApp(ir *gothic.Interpreter) {
    must(ir, `ttk::frame .appframe`)
    must(ir, `pack .appframe -fill both`)
    
    addTabbedPanel(ir, `.appframe`)
}

func must(ir *gothic.Interpreter, cmd string) {
    if err:= ir.Eval(cmd); err != nil {
        panic(err.Error())
    }
}

func addTabbedPanel(ir *gothic.Interpreter, appFrame string) {
    must(ir, `label .tabPanel -text "this is the label text"`)
    must(ir, fmt.Sprintf(`pack .tabPanel %s`, appFrame))
}
