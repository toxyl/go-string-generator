// This is a simple application that generates one line per argument.
// Each argument is parsed using the library and running the app
// multiple times with the same arguments will yield different results,
// except for the integer range token which will always produce the same
// sequence.
// The first argument must be the data directory used for [:file] tokens.
package main

import (
	"fmt"
	"os"

	gostringgenerator "github.com/toxyl/go-string-generator"
)

var DataDir string = ""

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Missing arguments!\nUsage: %s [pattern 1] .. [pattern N]\n", os.Args[0])
		return
	}
	gen := gostringgenerator.NewGenerator(
		DataDir, // location of data files to use for the [:file] token
		func(err error) { // function to handle errors reported by the file cache
			fmt.Printf("Encountered file cache error: %s", err.Error())
		},
	)
	for _, p := range os.Args[1:] {
		fmt.Printf("%s\n", gen.Generate(p))
	}
}
