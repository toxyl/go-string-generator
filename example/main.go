package main

import (
	"fmt"

	gostringgenerator "github.com/toxyl/go-string-generator"
)

var (
	gen = gostringgenerator.NewGeneratorSimple()
)

func print(pattern ...string) {
	for _, p := range pattern {
		fmt.Printf("%s\n", gen.Generate(p))
	}
}

func printExample(title, pattern string, repeats int) {
	fmt.Printf("%d %s\n", repeats, title)
	for repeats > 0 {
		print(pattern)
		repeats--
	}
	fmt.Println("")
}

func printExamples() {
	printExample("UUIDs", "[#UUID]", 10)
	printExample("Hashes (length = 10)", "[#10]", 10)
	printExample("Hashes (random length)", "[#[1-20]]", 10)
	printExample("Integer From Range", "[1-10]", 10)
	printExample("Integers (length = 10)", "[int:10]", 10)
	printExample("Integers (random length)", "[int:[1-20]]", 10)
	printExample("Strings (length = 10, lowercase)", "[str:10]", 10)
	printExample("Strings (length = 10, uppercase)", "[strU:10]", 10)
	printExample("Strings (length = 10, mixed case)", "[strR:10]", 10)
	printExample("Alphanumerics (length = 10, lowercase)", "[mix:10]", 10)
	printExample("Alphanumerics (length = 10, uppercase)", "[mixU:10]", 10)
	printExample("Alphanumerics (length = 10, mixed case)", "[mixR:10]", 10)
	printExample("List Of Integers", "[1..10]", 3)
	printExample("Strings From List", "[a,b,c,1,2,3]", 10)
	printExample("Strings From Random List", "[[str:1],[int:2],[mix:3]] - [#[[4..16],UUID]]", 10)
	printExample("Base64-Encoded Strings", "[b64:[str:20]]", 10)
	printExample("URL-Encoded Strings", "[url:[str:5]&[str:4]]", 10)
}

func main() {
	printExamples()
}
