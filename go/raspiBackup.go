package main

import (
	"fmt"
	"os"
	"path"
)

// Version -
const Version = "v0.7"

func main() {

	fmt.Printf("Hello %s %s\n", path.Base(os.Args[0]), Version)

}
