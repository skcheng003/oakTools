package main

import (
	"fmt"
	"github.com/oakTools/slice"
)

func main() {
	strs := []string{"hello", "my", "dear", "friend", "go", "lang"}

	strs, _= slice.Delete(strs, 1)
	fmt.Printf("The cap: %d\n", cap(strs))
	fmt.Printf("The nums %v\n", strs)

	strs, _ = slice.Delete(strs, 1)
	strs, _ = slice.Delete(strs, 1)
	strs, _ = slice.Delete(strs, 1)
	strs, _ = slice.Delete(strs, 1)

	fmt.Printf("The len: %d\n", len(strs))
	fmt.Printf("The cap: %d\n", cap(strs))
	fmt.Printf("The nums %v\n", strs)
}
