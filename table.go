package loopfmt

import "strings"

const (
	sep  = " | "
	line = "-"
)

type TableRow []string
type Table []TableRow

func NewTable() Table {
	return Table{}
}

func (t *Table) Append(items ...string) *Table {
	var row TableRow
	for i := range items {
		row = append(row, items[i])
	}
	*t = append(*t, row)
	return t
}

func (t *Table) AppendRows(rows ...TableRow) *Table {
	for i := range rows {
		*t = *t.Append(rows[i]...)
	}
	return t
}

func (t *Table) Render() string {
	if len(*t) == 0 {
		return ""
	}

	var res string
	res = sep + strings.Join((*t)[0], sep) + sep + "\n"
	res += sep + strings.Repeat(line+sep, len((*t)[0])) + "\n"

	for i := 1; i < len(*t); i++ {
		res += sep
		for j := range (*t)[i] {
			res += (*t)[i][j] + sep
		}
		res += "\n"
	}

	return res
}
