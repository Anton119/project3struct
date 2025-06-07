package main

import (
	"fmt"
	"project3struct/bins"
	"project3struct/config"
	"project3struct/storage"
	"time"

	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("ошибка с env")
	}

	cfg := config.NewConfig()
	fmt.Println("Key из когфига:", cfg)

	id := "1"
	private := true
	createdAt := time.Now()
	name := "first bin"
	FirstBin, _ := bins.NewBin(id, private, createdAt, name)

	id2 := "2"
	private2 := true
	createdAt2 := time.Now()
	name2 := "second bin"
	SecondBin, _ := bins.NewBin(id2, private2, createdAt2, name2)

	List := bins.NewBinList()
	List.AddBinToList(*FirstBin)
	List.AddBinToList(*SecondBin)
	fmt.Println(List)

	fmt.Println()
	fmt.Println("после преобразования в json")

	var store storage.Storage = &storage.FileStorage{Path: "data.json"}
	store.Save(List)
	fmt.Println(List)

	var storeTxt storage.Storage = &storage.FileStorage{Path: "store.txt"}
	storeTxt.Save(List)
	fmt.Println(List)

}

// для проверки
