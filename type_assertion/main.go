package main

import (
	"fmt"
	"reflect"
)

type ID string

type Utka struct {
	Name string
	Kys  []string
}

func main() {
	//Println3(123)
	//Println3("asd")
	//Println3(ID("123"))
	//typesLineChangingForEmptyInterface()
	//kindExample2()
	//
	//printStructInfo(Utka{
	//	Name: "Utka1",
	//	Kys:  []string{"eda"},
	//})

	printStructInfo(&Utka{
		Name: "Utka1",
		Kys:  []string{"eda"},
	})
}

func printStructInfo(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)
	fmt.Printf("-----Kind: %v-------\n", t.Kind())

	if t.Kind() == reflect.Pointer {
		v = reflect.Indirect(v) // fmt.Println(v.Interface()) -> will return empty interface value
		t = reflect.TypeOf(v.Interface())
	}

	// add check that if it is a pointer to a struct also proccess it
	if t.Kind() != reflect.Struct {
		fmt.Printf("ERR: Not a struct, expected struct value of 'kind' struct")
		return
	}

	n := t.NumField() // will panic if Kind is not struct
	fmt.Printf("Struct of type '%v' has %v fields\n", t, n)
	for i := 0; i < n; i++ {
		tt := t.Field(i)
		vv := v.Field(i)
		fmt.Printf("Field %v: name: %v, type: %v, value: %v\n\n", i, tt.Name, tt.Type, vv)
	}

	fmt.Println("---------")
}

func Println(x interface{}) {
	fmt.Printf("type: %T, value: %v\n", x, x)
}

func Println2(x interface{}) {
	if v, ok := x.(ID); ok {
		fmt.Printf("type: THIS IS ID, value: %v\n", v)
	} else {
		fmt.Printf("type: %T, value: %v\n", v, v)
	}
}

func Println3(x interface{}) {
	switch x.(type) {
	case ID:
		fmt.Println("ID!!!")
	default:
		fmt.Println("Not ID")
	}
}

func typesLineChangingForEmptyInterface() {
	// Выглядит как динамическое изменение типа, но на самом деле тип переменной остается interface{}
	// У которого есть два скрытых поля value, type и мы просто переписываем значения поля type для этого поля
	var x interface{}

	x = 3.14
	fmt.Printf("Type: %T, Value: %v", x, x)

	x = &struct{ name string }{}
	fmt.Printf("Type: %T, Value: %v", x, x)
}

func reflectExample() {
	x := 3.14

	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t)
	fmt.Println(v)
}

func kindExample() {
	var x interface{}

	x = new(string)

	t1 := reflect.TypeOf(x)
	v1 := reflect.ValueOf(x)

	fmt.Printf("super type: %v, type: %v, value: %v", t1.Kind(), t1, v1)

	a0 := [5]int{}
	a1 := [10]float64{}

	fmt.Println()
	// both of items "super types" is array
	fmt.Printf("kind1: %v, kind2: %v", reflect.TypeOf(a0).Kind(), reflect.TypeOf(a1).Kind())
}

func kindExample2() {
	var x interface{}

	x = "Hello"

	fmt.Println(reflect.TypeOf(x).Kind())

	type ID string

	x = ID("hello world")

	// anyway base type of ID is string
	fmt.Println(reflect.TypeOf(x).Kind())
}
