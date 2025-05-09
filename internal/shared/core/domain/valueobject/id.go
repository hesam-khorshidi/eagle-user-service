package valueobject

import (
	"strconv"
)

type ID int64

func (id ID) IsZero() bool {
	return id == 0
}

func (id ID) String() string {
	return strconv.FormatInt(int64(id), 10)
}
