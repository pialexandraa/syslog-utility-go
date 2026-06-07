package fileops

import (
	"bufio"
	"fmt"
	"os"
)

func CreateLogFile() {
	if !IsAdmin() {
		fmt.Println("Access denied: You do not have enough permisions to execute this action. Try again with Admin privileges.")
		return
	}

	fmt.Println("Enter the name of the log file you'd like to create:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()

	/* Creates the file here */
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	fmt.Printf("Log file '%s' created successfully.\n", fileName)
}

func ReadLogFile() {
	fmt.Println("Enter the name of the log file you'd like to read from:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()

	/* Reads the file here */
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Log file content:")
	fmt.Println(string(content))
}

func WriteLogEntry() {
	if !IsAdmin() {
		fmt.Println("Access denied: You do not have enough permisions to execute this action. Try again with Admin privileges.")
		return
	}

	fmt.Println("Enter the name of the log file you'd like to write to:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	fileName := scanner.Text()

	fmt.Println("Enter the log entry:")

	entryScanner := bufio.NewScanner(os.Stdin)
	entryScanner.Scan()
	data := entryScanner.Text()

	file, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening log file:", err)
		return
	}
	defer file.Close()

	/* Writes the string here */
	bytesWritten, err := file.WriteString(data + "\n")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Log entry successfully written to file. %d bytes written.\n", bytesWritten)
}
