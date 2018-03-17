package main

import (
	"fmt"
	"os"
	"path"
)

const Version = "v0.8"
const GitDate = "$Date: 2017-10-31 12:03:40 +0100$"
const GitSHA = "$Sha1: ed54e5e$"

func main() {

	fmt.Printf("%s\n", path.Base(os.Args[0]))

}
