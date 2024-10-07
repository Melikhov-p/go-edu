package GetFilesFromDirectoryRecursive

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("All files in directory:")
	PrintAllFiles(".")
	fmt.Println("Files with filtered path in directory:")
	PrintAllFilesWithFilter(".", "idea")
}

func PrintAllFiles(path string) {
	// получаем список всех элементов в папке (и файлов, и директорий)
	files, err := os.ReadDir(path)
	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}
	//  проходим по списку
	for _, f := range files {
		// получаем имя элемента
		// filepath.Join — функция, которая собирает путь к элементу с разделителями
		filename := filepath.Join(path, f.Name())
		// печатаем имя элемента
		fmt.Println(filename)
		// если элемент — директория, то вызываем для него рекурсивно ту же функцию
		if f.IsDir() {
			PrintAllFiles(filename)
		}
	}
}

func PrintAllFilesWithFilter(path string, filter string) {
	files, err := os.ReadDir(path)

	if err != nil {
		fmt.Println("unable to get list of files", err)
		return
	}

	for _, f := range files {
		filename := filepath.Join(path, f.Name())
		if strings.Contains(filename, filter) {
			fmt.Println(filename)
			// если элемент — директория, то вызываем для него рекурсивно ту же функцию
			if f.IsDir() {
				PrintAllFilesWithFilter(filename, filter)
			}
		}
	}
}

func PrintFilesWithFuncFilter(path string, predicate func(string) bool) {
	var walk func(string)
	walk = func(path string) {
		files, err := os.ReadDir(path)
		if err != nil {
			fmt.Println("unable to get list of files", err)
			return
		}
		for _, f := range files {
			filename := filepath.Join(path, f.Name())
			if predicate(filename) {
				fmt.Println(filename)
				if f.IsDir() {
					walk(filename)
				}
			}

		}
	}
	walk(path)
}

// containsDot возвращает все пути, содержащие точки
func containsDot(s string) bool {
	return strings.Contains(s, ".")
}
