package file

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ReadFile(name string) ([]byte, error) {
	data, err := os.ReadFile(name)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return data, nil

}

func WriteFile(content []byte, name string) {
	file, err := os.Create(name)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	_, err = file.Write(content)
	if err != nil {

		fmt.Println(err)
		return
	}
	fmt.Println("Запись успешна")

}

func IsJsonFile(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	return ext == ".json"
}
