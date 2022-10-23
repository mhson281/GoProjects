package processRequest

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func getRequest() []string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter a command and data: ")
	scanner.Scan()
	request := strings.Split(scanner.Text(), " ")

	return request
}

func ProcessRequest(dataFiles []string, maxFiles int) {
	for {
		request := getRequest()
		if len(request) < 1 {
			error := fmt.Errorf("[Error] Missing position argument")
			fmt.Println(error)
		}
		command := request[0]
		option := request[1:]
		switch command {
		case "exit":
			fmt.Println("[Info] Bye!")
			os.Exit(0)
		case "create":
			if len(option) < 1 {
				error := fmt.Errorf("[Error] Missing note argument")
				fmt.Println(error)
			} else if len(dataFiles) < maxFiles {
				addNote(&dataFiles, option)
			} else {
				error := fmt.Errorf("[Error] Notepad is full")
				fmt.Println(error)
			}
		case "clear":
			dataFiles = clearCache()
		case "list":
			if len(dataFiles) > 0 {
				printCache(&dataFiles)
			} else {
				error := fmt.Errorf("[Info] Notepad is empty")
				fmt.Println(error)
			}
		case "update":
			if len(option) == 0 {
				error := fmt.Errorf("[Error] Missing position argument")
				fmt.Println(error)
			} else {
				pos := option[0]
				note := option[1:]
				if len(note) < 1 {
					error := fmt.Errorf("[Error] Missing note argument")
					fmt.Println(error)
				} else {
					index, err := strconv.Atoi(pos)
					if err != nil {
						error := fmt.Errorf("[Error] Invalid position: %s", pos)
						fmt.Println(error)
						break
					} else if len(dataFiles) == 0 {
						error := fmt.Errorf("[Error] There is nothing to update")
						fmt.Println(error)
						break
					} else if index < 0 || index > maxFiles-1 {
						error := fmt.Errorf("[Error] Position %d is out of the boundary [1, %d]", index, maxFiles)
						fmt.Println(error)
						break
					}
					updateNote(&dataFiles, index, note, maxFiles)
				}
			}
		case "delete":
			if len(option) == 0 {
				error := fmt.Errorf("[Error] Missing position argument")
				fmt.Println(error)
			} else {
				pos := option[0]
				index, err := strconv.Atoi(pos)
				if err != nil {
					error := fmt.Errorf("[Error] Invalid position: %s", pos)
					fmt.Println(error)
					break
				} else if len(dataFiles) == 0 {
					error := fmt.Errorf("[Error] There is nothing to delete")
					fmt.Println(error)
					break
				} else if index < 0 || index > maxFiles-1 {
					error := fmt.Errorf("[Error] Position %d is out of the boundary [1, %d]", index, maxFiles)
					fmt.Println(error)
					break
				}
				dataFiles = deleteNote(&dataFiles, index, maxFiles)
				fmt.Printf("[OK] The note at position %d was successfully deleted\n", index)
			}

		default:
			fmt.Println("[Error] Unknown command")
		}
	}
}

func clearCache() []string {
	fmt.Println("[OK] All notes were successfully deleted")
	return make([]string, 0)
}

func addNote(files *[]string, command []string) {
	*files = append(*files, strings.Join(command, " "))
	fmt.Println("[OK] The note was successfully created")
}

func updateNote(files *[]string, index int, note []string, maxFiles int) {
	(*files)[index-1] = strings.Join(note, " ")
	fmt.Printf("[OK] The note at position %d was successfully updated\n", index)
}

func deleteNote(files *[]string, index int, maxFiles int) []string {
	newData := append((*files)[:index-1], (*files)[index:]...)
	return newData
}

func printCache(files *[]string) {
	for i, el := range *files {
		if len(el) > 0 {
			fmt.Printf("[Info] %d: %s\n", i+1, el)
		}
	}
}
