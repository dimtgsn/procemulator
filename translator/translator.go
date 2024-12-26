package translator

import (
	"procemulator/global"
	"strconv"
)

type Translator struct{}

func New() *Translator { return &Translator{} }

func (*Translator) Translate(memory []string) [1024]int {
	return translate(memory)
}

type Command string

const (
	PUSH  Command = "PUSH"
	READ  Command = "READ"
	STORE Command = "STORE"
	ADD   Command = "ADD"
	CMP   Command = "CMP"
	INC   Command = "INC"
	JMP   Command = "JMP"
	END   Command = "END"
)

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
		case string(PUSH):
			machineCode = 0b1
		case string(READ):
			machineCode = 0b10
		case string(STORE):
			machineCode = 0b11
		case string(ADD):
			machineCode = 0b100
		case string(CMP):
			machineCode = 0b101
		case string(INC):
			machineCode = 0b110
		case string(END):
			machineCode = 0b111
		case string(JMP):
			machineCode = 0b1000
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
