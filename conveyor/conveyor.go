package conveyor

import "fmt"

// Предобработчик для замены меток на адреса в памяти
type Preprocessor interface {
	PreprocessLabels(memory []string) []string
}

// Переводит данные в машинный код
type Translator interface {
	Translate(memory []string) [1024]int
}

// Интерпретирует машинный код
type Interpreter interface {
	Interpret(machineCode [1024]int) [1024]int
}

// Запускает процессор
func Run(
    preprocessor Preprocessor,
    translator Translator,
    interpreter Interpreter,
    memory []string,
) [1024]int {
	fmt.Println("START")

	preprocessedMemory := preprocessor.PreprocessLabels(memory)
    translatedMemory := translator.Translate(preprocessedMemory)
    endMemory := interpreter.Interpret(translatedMemory)
	fmt.Println("END")

	return endMemory
}