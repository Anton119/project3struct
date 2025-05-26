package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
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
		id:        id,
		private:   private,
		createdAt: createdAt,
		name:      name,
	}, nil
}
