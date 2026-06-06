package main

import (
	"bufio"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func createLogDirectory() {
	fmt.Println("Enter the name of the log directory you'd like to create:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	dirName := scanner.Text()

	/* Makes the directory here */
	err := os.Mkdir(dirName, 0755)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println("Successful creation of the directory: ", dirName)
}

func changeLogDirectory() {
	cwd := getWorkingDirectory()
	fmt.Println("Current directory is:", cwd)

	fmt.Print("Enter the name of the log directory you want to change to: ")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	newDirName := scanner.Text()

	/* Changes directory here */
	err := os.Chdir(newDirName)
	if err != nil {
		fmt.Println(err)
		return
	}

	newCwd := getWorkingDirectory()
	fmt.Println("Directory changed successfully. Current directory is:", newCwd)
}

func printWorkingDirectory() {
	cwd := getWorkingDirectory()
	fmt.Println("Current working directory is:", cwd)
}

func getWorkingDirectory() string {
	/* Gets the working directory here */
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return cwd
}

func listDirectoryContents() {
	cwd := getWorkingDirectory()
	fmt.Println("Current Working directory is: ", cwd)

	files, err := os.ReadDir(cwd)
	if err != nil {
		fmt.Println("Error reading directory contents:", err)
		return
	}

	fmt.Println("Directory contents:")
	for _, file := range files {
		fmt.Println(file.Name())
	}
}

func listDirectoryContentsRecursive() {
	if !isAdmin() {
		fmt.Println("The access is denied! Note: this is an Admin view. Please run the utility as admin (e.g.'with sudo') or ask for admin help.")
		return
	}
	cwd := getWorkingDirectory()
	fmt.Println("The starting point/current working directory is: ", cwd)

	err := filepath.WalkDir(cwd, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return nil // silently skips directories that even the admin does not have enough priviledges to access
		}
		fmt.Println(path)
		return nil
	})

	if err != nil {
		fmt.Println("Error reading the directory recursively: ", err)
	}
	fmt.Println("Finalized recursive data checks.")

}
