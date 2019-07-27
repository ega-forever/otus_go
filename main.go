package main

import (
	"errors"
	"fmt"
	"reflect"
	"unicode"
	"unsafe"
)

func structToMap(iv interface{}) (map[string]interface{}, error) {
	v := reflect.ValueOf(iv)
	if v.Kind() != reflect.Struct {
		return nil, errors.New("not a struct")
	}
	t := v.Type()
	mp := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i) // reflect.StructField
		fv := v.Field(i)    // reflect.Value

		if unicode.IsUpper([]rune(field.Name)[0]) == false {
			continue
		}

		mp[field.Name] = fv.Interface()
	}
	return mp, nil
}

func main() {

	// type assertion
	/*	var r io.Reader
		var f *os.File
		var ok bool
		f, ok = r.(*os.File)
	*/

	//var i int = 42
	var s = struct {
		d string
		S int
	}{"hello", 42}
	fmt.Println(s)

	sv := reflect.ValueOf(&s)
	fmt.Println(sv.Kind())
	fmt.Println(sv.Elem().NumField())
	fmt.Println(structToMap(s))

	reflFunc := reflect.ValueOf(structToMap)

	args := []reflect.Value{reflect.ValueOf(s)}

	fmt.Println(reflFunc.Call(args))

	type St struct{ a, b int32 }
	var b [8]byte
	bp := &b
	var sp *St
	var up unsafe.Pointer
	up = unsafe.Pointer(bp)
	sp = (*St)(up)
	sp.a = 12345678
	fmt.Println(b)

}
