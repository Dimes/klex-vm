package main

import (
	"fmt"
	"strconv"
	"secd"
	"secd/command"
	"secd/runner"
)

func main() {
	machine := secd.NewSecd()

	rootEnv := secd.NewEnvironment()
	machine.E.Push(rootEnv)

	// Machine code for:
	// func f(a, b) {
	//     return a - b
	// }
	// const a = 40
	// f(a, 5)

	commands := []command.Command {
		command.LoadFunc(2, []command.Command {
			command.Load(1, 1), // Push b
			command.Load(1, 0), // Push a
			command.SubInt(),   // a - b
			command.Return(),
		}),
		command.InitVar(),    // (0, 0)
		command.LoadConst(40),
		command.InitVar(),    // (0, 1)
		command.LoadConst(5), // Push b (5)
		command.Load(0, 1),   // Push a (40)
		command.Load(0, 0),   // Push f
		command.Apply(),      // f(a, b)
		command.Stop(),
	}

	for i := len(commands) - 1; i >= 0; i-- {
		machine.C.Push(commands[i])
	}

	for !machine.IsStopped() {
		runner.Step(machine)
	}
	
	fmt.Printf("Program complete\n")
	fmt.Printf("Final value: " + strconv.Itoa(machine.S.Pop().(int)) + "\n")
}
