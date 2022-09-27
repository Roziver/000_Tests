package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Введите длину: ")

	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	size, err := strconv.ParseFloat(input, 64)
	if err != nil {
		log.Fatal(err)
	}
	var status string
	if size <= 14 {
		status = "Печаль"
	}
	if size == 15 {
		status = "Не приговор"
	}
	if size >= 16 {
		status = "Норм"
	}
	fmt.Println(status)
}
