package translator

import (
	"procemulator/global"
	"strconv"
)

type Translator struct {}

func New() *Translator { return &Translator{} }

func (*Translator) Translate(memory []string) [1024]int {
	return translate(memory)
}

func translate(memory []string) [1024]int {
	var machineMemory [1024]int

	for idx, value := range memory[:global.MEMORYSIZE] {
		if value != "" {
			if intValue, err := strconv.Atoi(value); err == nil {
				machineMemory[idx] = intValue
			}
		}
	}

	// Команды в машинный код
	for idx, command := range memory[global.MEMORYSIZE:] {
		if command == "" {
			break
		}
		var machineCode int

		switch command {
		case "PUSH":
			machineCode = 0b1
		case "READ":
			machineCode = 0b10
		case "WRITE":
			machineCode = 0b11
		case "ADD":
			machineCode = 0b100
		case "DUP":
			machineCode = 0b101
		case "DROP":
			machineCode = 0b110
		case "LDC":
			machineCode = 0b111
		case "STC":
			machineCode = 0b1000
		case "CMP":
			machineCode = 0b1001
		case "SWAP":
			machineCode = 0b1010
		case "INC":
			machineCode = 0b1011
		case "END":
			machineCode = 0b1100
		case "JUMP":
			machineCode = 0b1111
		default:
			if intValue, err := strconv.Atoi(command); err == nil {
				machineCode = intValue
			}
		}

		// проверяем карту меток
		if i, exists := global.LABELS[command]; exists {
			machineMemory[idx+global.MEMORYSIZE] = i
		} else {
			machineMemory[idx+global.MEMORYSIZE] = machineCode
		}
	}

	return machineMemory
}