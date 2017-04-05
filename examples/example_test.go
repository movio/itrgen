package main

import (
	"fmt"
	"testing"
)

func TestStringItr(t *testing.T) {
	list := StringItr{
		"John",
		"Jane",
	}

	res, _ := list.MapString(func(i string) (string, error) {
		return fmt.Sprintf("Hello %s", i), nil
	})

	expected := []string{
		"Hello John",
		"Hello Jane",
	}

	for i := 0; i < 2; i++ {
		if expected[i] != res[i] {
			t.Fatalf("want %v, have %v", expected[i], res[i])
		}
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

	expected := []string{
		"Hi Sam",
		"Hi Sarah",
	}

	for i := 0; i < 2; i++ {
		if expected[i] != res[i] {
			t.Fatalf("want %v, have %v", expected[i], res[i])
		}
	}
}
