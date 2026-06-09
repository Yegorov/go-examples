//go:build ignore

// go version
// go run cmd/share/unique.go
package main

import (
	"fmt"
	"unique"
	"unsafe"
)

func main() {
	str0 := "Go"
	str1 := *new("Go")
	str2 := *new("Go")
	str3 := unique.Make("Go").Value()
	str4 := unique.Make("Go").Value()
	fmt.Printf(
		"str0: %p, str1: %p, str2: %p, str3: %p, str4: %p\n",
		&str0, &str1, &str2, &str3, &str4,
	)
	fmt.Printf(
		"str0: %p, str1: %p, str2: %p, str3: %p, str4: %p\n",
		unsafe.StringData(str0),
		unsafe.StringData(str1),
		unsafe.StringData(str2),
		unsafe.StringData(str3),
		unsafe.StringData(str4),
	)
}
