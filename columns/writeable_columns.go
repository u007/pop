package columns

import (
	"sort"
	"strings"
)

type WriteableColumns struct {
	Columns
}

func (c WriteableColumns) UpdateString() string {
	xs := []string{}
	for _, t := range c.Cols {
		xs = append(xs, t.UpdateString())
	}
	if SortColumns {
		sort.Strings(xs)
	}
	return strings.Join(xs, ", ")
}
