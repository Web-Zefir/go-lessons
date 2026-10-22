package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func printMainMenu(folders []string) {
	fmt.Println("=== папки ===")
	for i, name := range folders {
		fmt.Printf("%d. %s\n", i+1, name)
	}
	fmt.Println("0. Выход")
	fmt.Print("> ")
}

func createFolder(folders []string, name string) ([]string, error) {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return folders, errors.New("имя папки не может быть пустым")
	}

	for _, f := range folders {
		if f == trimmed {
			return folders, errors.New("папка с таким именем уже существует")
		}
	}

	return append(folders, trimmed), nil
}

func deleteFolder(folders []string, name string) ([]string, error) {
	index := -1
	for i, f := range folders {
		if f == name {
			index = i
			break
		}
	}

	if index == -1 {
		return folders, errors.New("папки нет")
	}

	return append(folders[:index], folders[index+1:]...), nil
}

func renameFolder(folders []string, oldName, newName string) ([]string, error) {
	trimmedNew := strings.TrimSpace(newName)
	if trimmedNew == "" {
		return folders, errors.New("новое имя папки не может быть пустым")
	}

	index := -1
	for i, f := range folders {
		if f == oldName {
			index = i
		}
		if f == trimmedNew {
			return folders, errors.New("папка с таким именем уже существует")
		}
	}

	if index == -1 {
		return folders, errors.New("папки нет")
	}

	folders[index] = trimmedNew
	return folders, nil
}

func openFolderMenu(currentFolder *string, folders *[]string, scanner *bufio.Scanner) bool {
	for {
		fmt.Printf("\n[%s]\n", *currentFolder)
		fmt.Println("задач пока нет — занятие 02")
		fmt.Println("1. Создать папку")
		fmt.Println("2. Удалить папку")
		fmt.Println("3. Переименовать папку")
		fmt.Println("4. Назад")
		fmt.Println("0. Выход")
		fmt.Print("> ")

		if !scanner.Scan() {
			return true
		}
		cmd := strings.TrimSpace(scanner.Text())

		switch cmd {
		case "1":
			fmt.Print("Введите имя новой папки: ")
			if scanner.Scan() {
				newName := scanner.Text()
				updated, err := createFolder(*folders, newName)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					*folders = updated
					fmt.Printf("Папка '%s' успешно создана!\n", strings.TrimSpace(newName))
				}
			}

		case "2":
			updated, err := deleteFolder(*folders, *currentFolder)
			if err != nil {
				fmt.Printf("Ошибка: %v\n", err)
			} else {
				*folders = updated
				fmt.Printf("Папка '%s' удалена.\n", *currentFolder)
				return false
			}

		case "3":
			fmt.Print("Введите новое имя папки: ")
			if scanner.Scan() {
				newName := scanner.Text()
				updated, err := renameFolder(*folders, *currentFolder, newName)
				if err != nil {
					fmt.Printf("Ошибка: %v\n", err)
				} else {
					*folders = updated
					*currentFolder = strings.TrimSpace(newName)
					fmt.Printf("Папка успешно переименована в '%s'!\n", *currentFolder)
				}
			}

		case "4":
			return false

		case "0":
			return true

		default:
			fmt.Println("Неизвестная команда, попробуйте еще раз.")
		}
	}
}

func main() {
	folders := []string{"Входящие", "Работа", "Личное", "Архив"}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		printMainMenu(folders)

		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		if choice == "0" {
			fmt.Println("Пока!")
			return
		}

		index, err := strconv.Atoi(choice)
		if err != nil || index < 1 || index > len(folders) {
			fmt.Println("Ошибка: неверный пункт меню!\n")
			continue
		}

		currentFolder := folders[index-1]

		shouldExit := openFolderMenu(&currentFolder, &folders, scanner)
		if shouldExit {
			fmt.Println("Пока!")
			return
		}

		fmt.Println()
	}
}
