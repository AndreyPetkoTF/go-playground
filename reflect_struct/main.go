package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	res := dataToStruct([]TableField{
		{Name: "ID", Type: "int", Value: "1"},
		{Name: "Name", Type: "string", Value: "James"},
	})

	fmt.Printf("%#v", res)
}

type TableField struct {
	Value string
	Name  string
	Type  string
}

func dataToStruct(data []TableField) interface{} {
	structFields := make([]reflect.StructField, len(data))

	var t reflect.Type

	for i, tf := range data {
		structFields[i].Name = tf.Name

		switch tf.Type {
		case "int":
			t = reflect.TypeOf(0)
		case "uint8":
			t = reflect.TypeOf(uint8(0))
		case "string":
			t = reflect.TypeOf("")
		case "float64":
			t = reflect.TypeOf(0.0)
		}
		structFields[i].Type = t
	}

	typ := reflect.StructOf(structFields)
	v := reflect.New(typ).Elem()
	s := v.Addr().Interface()

	for i, tf := range data {
		switch tf.Type {
		case "int":
			tv, _ := strconv.ParseInt(tf.Value, 10, 64)
			v.Field(i).SetInt(tv)
		case "string":
			v.Field(i).SetString(tf.Value)
		}
	}

	return s
}
