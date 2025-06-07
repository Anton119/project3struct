package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"project3struct/bins"
	"project3struct/file"
)

type Storage interface {
	Save(list *bins.BinList) error
}

type FileStorage struct {
	Path string
}

func (s *FileStorage) Save(list *bins.BinList) error {
	data, err := ListToBytes(list)
	if err != nil {
		fmt.Println("не удалось преобразовать")
		return err
	}

	err = file.WriteFile(data, s.Path)
	if err != nil {
		fmt.Println("не удалось записать файл")
		return err
	}
	return nil
}

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
