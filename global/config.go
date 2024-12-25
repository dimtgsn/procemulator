package global

var (
	// Карта для хранения адресов меток в памяти
	LABELS      = make(map[string]int)

	// Размер памяти с инструкциями
	PROGRAMSIZE = 896

	// Размер памяти с данными
	MEMORYSIZE  = 128
	
	// Размер одной команды в памяти
	COMMANDSIZE = 16
	
	// Массив памяти
	Memory      [1024]string
	
	// Указатель стека
	SC          = 128
	
	// Указатель программы
	PC          = MEMORYSIZE
)

func StackIsEmpty() bool {
	return SC >= 128
}