package loopfmt

import "strings"

type List struct {
	data   []string
	prefix string
}

func NewOrderedList() List {
	return List{
		prefix: " 1. ",
	}
}

func NewUnorderedList() List {
	return List{
		prefix: " - ",
	}
}

func (l *List) Append(items ...string) *List {
	for i := range items {
		l.data = append(l.data, items[i])
	}
	return l
}

func (l *List) Render() string {
	return strings.Join(mapf(l.data, MapPrefix(l.prefix)), "\n")
}

type MapFunc func(string) string

func mapf(items []string, mapFunc ...MapFunc) []string {
	var res []string
	for i := range items {
		r := items[i]
		for j := range mapFunc {
			r = mapFunc[j](r)
		}
		res = append(res, r)
	}
	return res
}

func MapPrefix(prefix string) MapFunc {
	return func(s string) string {
		return prefix + s
	}
}
