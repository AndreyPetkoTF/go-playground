package main

import "fmt"

type Krea struct {
	ID int32
}

func main() {
	items := []Item1{
		{ID: 1, name: "2"},
	}

	res := ToMap(items)

	fmt.Printf("%v %T\n", res, res)

	kreas := []Krea{
		{ID: 1},
		{ID: 2},
	}

	var skills = ToAliasMap[Krea](kreas, func(k Krea) Alias {
		return Alias{
			ID:    k.ID,
			Alias: "a",
			Name:  "b",
		}
	})

	fmt.Println("hello")
	fmt.Println(skills)
}

type Item1 struct {
	ID   int
	name string
}

func (i Item1) GetId() int {
	return i.ID
}

type Item2 struct {
	ID    int
	name2 string
}

func (i Item2) GetId() int {
	return i.ID
}

type GeneralItem interface {
	GetId() int
}

func ToMap[T GeneralItem](items []T) map[int]T {
	mapRes := make(map[int]T, len(items))
	for _, item := range items {
		mapRes[item.GetId()] = item
	}

	return mapRes
}
