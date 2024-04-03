// This is a simple application that generates one line per argument.
// Each argument is parsed using the library and running the app
// multiple times with the same arguments will yield different results,
// except for the integer range token which will always produce the same
// sequence.
// [:file] tokens will be looked up from $HOME/.rsg-data
package main

import (
	"fmt"
	"os"

	gostringgenerator "github.com/toxyl/go-string-generator"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Missing arguments!\nUsage: %s [pattern 1] .. [pattern N]\n", os.Args[0])
		return
	}

	gen := gostringgenerator.NewGeneratorSimple()
	for _, p := range os.Args[1:] {
		fmt.Printf("%s\n", gen.Generate(p))
	}
}
