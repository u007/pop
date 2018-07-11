package columns

import (
	"sort"
	"strings"
)

// ReadableColumns represents a list of columns Pop is allowed to read.
type ReadableColumns struct {
	Columns
}

// SelectString returns the SQL column list part of the SELECT
// query.
func (c ReadableColumns) SelectString() string {
	xs := []string{}

	for i := range c.ColSequence {
		name := c.ColSequence[i]

		if col, ok := c.Cols[name]; ok {
			xs = append(xs, col.SelectSQL)
		}
	}

	if SortColumns {
		sort.Strings(xs)
	}
	return strings.Join(xs, ", ")
}
