package main

import (
	"fmt"
	"reflect"
)

func main() {
	type T struct {
		a string `ct:"ct-a"`
		b int    `ct:"ct-b"`
	}

	t := T{a: "asd", b: 0}

	res := reflect.ValueOf(t)

	fmt.Println("can set ", res.CanSet())
	fmt.Println("can set ", res)

	fc := res.NumField()
	for i := 0; i < fc; i++ {
		f := res.Field(i)
		fmt.Println(f.Type().Name(), ": ", f)
	}

}
