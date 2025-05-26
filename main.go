package main

import (
	"fmt"
	"project3struct/bins"
	"time"
)

func main() {
	id := "bin123"
	private := true
	createdAt := time.Now()
	name := "My Bin"

	newBin, err := bins.NewBin(id, private, createdAt, name)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Bin created:", newBin)
}
