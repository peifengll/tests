package main

import (
	"fmt"
	"testing"
)

func TestGetNilSlice(t *testing.T) {
	oidPgid := make(map[string][]string)
	c := oidPgid["123"]
	fmt.Println(c == nil)
	c = make([]string, 0)
	oidPgid["123"] = c
	c = oidPgid["123"]
	fmt.Println(c == nil)
}

func TestNilSliceAppend(t *testing.T) {
	oidPgid := make(map[string][]string)
	c := oidPgid["123"]
	c = append(c, "123")
	fmt.Println(oidPgid["123"])
	oidPgid["123"] = c
	fmt.Println(oidPgid["123"])
}

func TestRangeKey(t *testing.T) {
	appendGids := make(map[string]bool)
	if appendGids["66"] == false {
		fmt.Println("000")
	} else {
		fmt.Println("111")
	}
	for k, v := range appendGids {
		fmt.Println(k, ":", v)
	}
}
