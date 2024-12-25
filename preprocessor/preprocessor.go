package preprocessor

import "procemulator/global"

type Preprocessor struct {}

func New() *Preprocessor { return &Preprocessor{} }

func (*Preprocessor) PreprocessLabels(memory []string) []string {
	const label = ':'
	return preprocess(memory, label)
}

func preprocess(memory []string, label byte) []string {
	var (
		processed []string
		shift int
	)

	for idx, command := range memory[global.MEMORYSIZE:] {
		if command == "" {
			break
		}

		if commandIsLabel(command, label) {
			global.LABELS[command[:len(command)-1]] = idx + global.MEMORYSIZE - shift
			shift++
		}
	}

	for _, command := range memory {
		if command == "" {
			break
		}
		if !commandIsLabel(command, label) {
			processed = append(processed, command)
		}
	}

	return processed
}

func commandIsLabel(command string, label byte) bool {
	return command[len(command)-1] == label
}
