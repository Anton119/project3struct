package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

type BinList struct {
	Bins []Bin
}

func NewBin(id string, private bool, createdAt time.Time, name string) (*Bin, error) {

	if id == "" {
		return nil, errors.New("INVALID_ID")
	}

	if name == "" {
		return nil, errors.New("INVALID_NAME")
	}
	return &Bin{
		Id:        id,
		Private:   private,
		CreatedAt: createdAt,
		Name:      name,
	}, nil
}

func NewBinList() *BinList {
	return &BinList{}

}

func (list *BinList) AddBinToList(bin Bin) error {

	if list == nil {

		return errors.New("EMPTY LIST")
	}

	list.Bins = append(list.Bins, bin)
	return nil
}
