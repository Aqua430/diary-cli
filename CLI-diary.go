package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Укажите команду: write, read, clear или exit")
		return
	}
	command := os.Args[1]

	switch command {
	case "write":
		writeEntry()
	case "read":
		readEntries()
	case "clear":
		clearDiary()
	case "exit":
		fmt.Println("Завершение программы.")
		return
	default:
		fmt.Println("Неизвестная команда:", command)
		fmt.Println("Введите write, read, clear или exit")
	}
}

func writeEntry() {
	file, err := os.OpenFile("diary.txt", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	fmt.Print("Введите новую запись: ")
	scanner := bufio.NewScanner(os.Stdin)

	if scanner.Scan() {
		now := time.Now()
		timestamp := now.Format("2006-01-02 15:04:05")
		line := scanner.Text()
		fullLine := timestamp + " — " + line + "\n"
		writer.WriteString(fullLine)
		writer.Flush()
		file.Sync()
		fmt.Println("Запись сохранена")
	} else if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Ошибка ввода: ", err)
		os.Exit(1)
	}
}

func readEntries() {
	file, err := os.Open("diary.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	fmt.Println("Содержимое дневника:")
	hasContent := false
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		hasContent = true
	}
	if !hasContent {
		fmt.Println("Дневник пуст.")
	}
}

func clearDiary() {
	err := os.WriteFile("diary.txt", []byte{}, 0644)
	if err != nil {
		panic(err)
	}
	fmt.Println("Содержимое дневника было очищено.")
}
