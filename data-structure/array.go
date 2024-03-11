package data

import "strings"

type Array struct {
	Elems []string
}

func (a *Array) InsertArray(data string) {
	if data == "" {
		println("data null")
	}

	a.Elems = append(a.Elems, data)
}

func (a *Array) DeletedArray(index int) string {
	if len(a.Elems) == 0 {
		return "sorry data is null"
	}

	a.Elems = append(a.Elems[:index], a.Elems[index+1:]...)

	return "success deleted"
}

func (a *Array) JoinArray(sep string) string {
	data := strings.Join(a.Elems, sep)
	return data
}
