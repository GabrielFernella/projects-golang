package main

import (
	"fmt"
	"os"
)

func ReadFile() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado")
		}
	}()

	_, err := os.Open("./settings.txt")
	if err != nil {
		panic("File not exists")
	}
}

func main() {

	ReadFile()

	fmt.Println("executed")
}

// defer serve para executar por ultimo no final da funcão
// recover geralmente é utilizado para algumas tratativas especídicas de exceções
