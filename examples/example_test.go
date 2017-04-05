package main

import (
	"fmt"
	"testing"
)

type MyStruct struct {
	Name string
}

//go:generate itrgen -type=MyStruct -map-to="MyStruct,string,int64"
//go:generate itrgen -package="main" -type="string" -map-to="MyStruct,string,int64"

func TestStringItr(t *testing.T) {
	list := StringItr{
		"John",
		"Jane",
	}

	res, _ := list.MapString(func(i string) (string, error) {
		return fmt.Sprintf("Hello %s", i), nil
	})

	res0 := "Hello John"
	if res[0] != res0 {
		t.Fatalf("want %v, have %v", res0, res[0])
	}
	res1 := "Hello Jane"
	if res[1] != res1 {
		t.Fatalf("want %v, have %v", res1, res[1])
	}

}

func TestStructItr(t *testing.T) {
	list := MyStructItr{
		MyStruct{"Sam"},
		MyStruct{"Sarah"},
	}

	res, _ := list.MapString(func(i MyStruct) (string, error) {
		return fmt.Sprintf("Hi %s", i.Name), nil
	})

	res0 := "Hi Sam"
	if res[0] != res0 {
		t.Fatalf("want %v, have %v", res0, res[0])
	}
	res1 := "Hi Sarah"
	if res[1] != res1 {
		t.Fatalf("want %v, have %v", res1, res[1])
	}
}
