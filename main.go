package main

import (
	"bufio"
	"fmt"
	"os"
	"procemulator/conveyor"
	"procemulator/global"
	"procemulator/interpreter"
	"procemulator/preprocessor"
	"procemulator/translator"
)

func main() {
	file, err := os.Open("memory-v2.txt")
	if err != nil {
		fmt.Println("failed to open file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	
	var (
		idx int
		memory [1024]string
	)
	for scanner.Scan() {
		memory[idx] = scanner.Text()
		idx++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("failed to read file:", err)
		return
	}
	
	endMemory := conveyor.Run(
		preprocessor.New(),
		translator.New(),
        interpreter.New(),
        memory[:],
	)
	fmt.Printf("Sum of array elements: %d\n", endMemory[global.SC])
}
