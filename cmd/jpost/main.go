package main

import (
	"fmt"
	"os"

	"github.com/syumai/go-jpostcode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "post code name is required. e.g. 0010010")
		return
	}
	address, err := jpostcode.Find(os.Args[1])
	if err == jpostcode.ErrNotFound {
		fmt.Fprintln(os.Stderr, "address was not found")
		return
	}
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %v\n", err)
		return
	}
	addressJSON, err := address.ToJSON()
	if err != nil {
		fmt.Fprintf(os.Stderr, "unexpected error: %v\n", err)
		return
	}
	fmt.Println(addressJSON)
}
