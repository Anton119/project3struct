package main

import (
	"fmt"
	"project3struct/bins"
	"project3struct/storage"
	"time"
)

func main() {

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

	storage.Save(List)
	fmt.Println(List)

}

// для проверки
