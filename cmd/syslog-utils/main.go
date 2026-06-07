package main

import (
	"fmt"
	"strconv"
	"syslog-utility-go/internal/fileops"
)

func main() {
	fmt.Println("The System Log utility is running. What would you like to do next?")

	for {
		printOptions()

		var input string
		fmt.Scanln(&input)

		if input == "exit" {
			break
		}

		// Convert the user input from a string to an integer
		option, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid option provided.")
			continue
		}

		switch option {
		case 1:
			fileops.CreateLogDirectory()
		case 2:
			fileops.ChangeLogDirectory()
		case 3:
			fileops.CreateLogFile()
		case 4:
			fileops.ReadLogFile()
		case 5:
			fileops.WriteLogEntry()
		case 6:
			fileops.PrintWorkingDirectory()
		case 7:
			fileops.ListDirectoryContents()
		case 8:
			fileops.ListDirectoryContentsRecursive()
		default:
			fmt.Println("Invalid option provided.")
		}
	}
}

func printOptions() {
	fmt.Println("--------------")
	fmt.Println("Please input a number to select an option or type 'exit' to quit:")
	fmt.Println("1 - Create a log directory")
	fmt.Println("2 - Change to a log directory")
	fmt.Println("3 - Create a log file")
	fmt.Println("4 - Read a log file")
	fmt.Println("5 - Write a log entry")
	fmt.Println("6 - Print working directory (pwd)")
	fmt.Println("7 - List directory contents (ls)")
	fmt.Println("8 - List the directory tree recursively (ls -R)")
}
