package interpreter

import (
	"procemulator/global"
)

type Interpreter struct {}

func New() *Interpreter { return &Interpreter{} }

func (*Interpreter) Interpret(machineCode [1024]int) [1024]int {
    return interpret(machineCode)
}

func interpret(machineCode [1024]int) [1024]int {
loop:
	for {
		curCommand := machineCode[global.PC]

		switch curCommand {
		case 0b0: // SKIP
			global.PC++
		case 0b1: // PUSH
			global.SC--
			global.PC++
			machineCode[global.SC] = machineCode[global.PC]
			global.PC++
		case 0b10: // READ
			if global.StackIsEmpty() {
				break loop
			}
			machineCode[global.SC] = machineCode[machineCode[global.SC]]
			global.PC++
		case 0b11: // WRITE
			if global.StackIsEmpty() {
				break loop
			}
			machineCode[machineCode[global.SC]] = machineCode[global.SC+1]
			global.SC += 2
			global.PC++
		case 0b1010: // SWAP
			if global.StackIsEmpty() {
				break loop
			}
			machineCode[global.SC], machineCode[global.SC+1] = machineCode[global.SC+1], machineCode[global.SC]
			global.PC++
		case 0b1011: // INC
			if global.StackIsEmpty() {
				break loop
			}
			machineCode[global.SC] = 1 + machineCode[global.SC]
			global.PC++
		case 0b1111: // JUMP
			if global.StackIsEmpty() {
				break loop
			}

			switch {
			case machineCode[global.SC+1] <= 0:
				machineCode[global.SC+1] = machineCode[global.SC]
				global.SC++
				global.PC = machineCode[global.SC]
			default:
				global.PC++
			}

		case 0b1001: // CMP
			if global.StackIsEmpty() {
				break loop
			}
			if (machineCode[global.SC] - machineCode[global.SC+1]) == 0 {
				machineCode[global.SC+1] = 0
			} else if (machineCode[global.SC] - machineCode[global.SC+1]) < 0 {
				machineCode[global.SC+1] = -1
			} else {
				machineCode[global.SC+1] = 1
			}
			global.SC++
			global.PC++
		case 0b1000: // STC:
			if global.StackIsEmpty() {
				break loop
			}
			global.SC--
			machineCode[global.SC] = global.PC
		case 0b110: // DROP
			if global.StackIsEmpty() {
				break loop
			}
			global.SC++
			global.PC++
		case 0b101: // DUP
			if global.StackIsEmpty() {
				break loop
			}
			global.SC--
			machineCode[global.SC] = machineCode[global.SC+1]
			global.PC++
		case 0b100: // ADD
			if global.StackIsEmpty() {
				break loop
			}
			global.SC++
			machineCode[global.SC] = machineCode[global.SC-1] + machineCode[global.SC]
			global.PC++
		case 0b1100: // END
			break loop
		}
	}

	return machineCode
}
