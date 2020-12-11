package puzzles

import (
	"errors"
	"io"
)

type handheldHalting struct {
}

type consoleVM struct {
	handlers     map[instructionType]instructionHandler
	instructions []instruction
	lastExecuted int
	curr         int
	acc          int
	executed     map[int]bool
}

type instructionHandler func(vm *consoleVM, arg int)

type instruction struct {
	kind     instructionType
	argument int
}

type instructionType string

const (
	acc instructionType = "acc"
	jmp instructionType = "jmp"
	nop instructionType = "nop"
)

var infiniteLoopError = errors.New("Infinite Loop Detected")

func NewConsoleVM(instructions []instruction) *consoleVM {
	c := new(consoleVM)
	c.instructions = instructions
	c.executed = make(map[int]bool)
	c.registerHandlers()
	return c
}

func (c *consoleVM) registerHandlers() {
	c.handlers = make(map[instructionType]instructionHandler)
	c.handlers[acc] = accHandler
	c.handlers[jmp] = jmpHandler
	c.handlers[nop] = nopHandler
}

func (h handheldHalting) Puzzle1(reader io.Reader) (Result, error) {
	instructions, err := ParseBootCode(reader)
	if err != nil {
		return nil, err
	}

	vm := NewConsoleVM(instructions)

	err = vm.run()
	if err == infiniteLoopError {
		return intResult(vm.acc), nil
	}

	return nil, errors.New("Expected Infinite Loop")

}

func (h handheldHalting) Puzzle2(reader io.Reader) (Result, error) {
	instructions, err := ParseBootCode(reader)
	if err != nil {
		return nil, err
	}
	var nopJmpSet = make(map[int]instructionType)
	for i, ins := range instructions {
		if ins.kind == nop || ins.kind == jmp {
			nopJmpSet[i] = ins.kind
		}
	}
	instructionsCopy := make([]instruction, len(instructions))
	var vm *consoleVM
	for k, v := range nopJmpSet {
		copy(instructionsCopy, instructions)
		if v == nop {
			instructionsCopy[k] = instruction{jmp, instructions[k].argument}
		}
		if v == jmp {
			instructionsCopy[k] = instruction{nop, instructions[k].argument}
		}

		vm = NewConsoleVM(instructionsCopy)
		err = vm.run()
		if err == nil {
			break
		}
	}

	return intResult(vm.acc), nil
}

func (c *consoleVM) run() error {
	for {
		if c.curr == len(c.instructions) {
			return nil
		}

		err := c.processNextInstruction()
		if err != nil {
			return err
		}
	}
}

func (c *consoleVM) processNextInstruction() error {
	if c.executed[c.curr] {
		return infiniteLoopError
	}

	ins := c.instructions[c.curr]
	c.lastExecuted = c.curr
	c.executed[c.curr] = true

	handler := c.handlers[ins.kind]
	handler(c, ins.argument)

	return nil
}

// Instruction Handlers

func accHandler(vm *consoleVM, arg int) {
	vm.acc += arg
	vm.curr++
}

func jmpHandler(vm *consoleVM, arg int) {
	vm.curr = vm.curr + arg
}

func nopHandler(vm *consoleVM, arg int) {
	vm.curr++
}