package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"project3struct/bins"
	"project3struct/file"
)

func BinToBytes(b *bins.Bin) ([]byte, error) {

	if b == nil {
		return nil, errors.New("bin is nil")
	}
	file, err := json.Marshal(b)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func ListToBytes(list *bins.BinList) ([]byte, error) {
	file, err := json.Marshal(list)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func Save(list *bins.BinList) {
	data, err := ListToBytes(list)
	if err != nil {
		fmt.Println("не удалось преобразовать")
	}

	file.WriteFile(data, "data.json")
}
