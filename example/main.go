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
	fmt.Printf("%d %s\nPattern: %s\n\n", repeats, title, pattern)
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
	printExample("Integers (length = 10, short syntax)", "[i:10]", 10)
	printExample("Strings (length = 10, lowercase)", "[str:10]", 10)
	printExample("Strings (length = 10, uppercase)", "[strU:10]", 10)
	printExample("Strings (length = 10, mixed case)", "[strR:10]", 10)
	printExample("Strings (length = 10, lowercase, short syntax)", "[s:10]", 10)
	printExample("Strings (length = 10, uppercase, short syntax)", "[sU:10]", 10)
	printExample("Strings (length = 10, mixed case, short syntax)", "[sR:10]", 10)
	printExample("Alphanumerics (length = 10, lowercase)", "[mix:10]", 10)
	printExample("Alphanumerics (length = 10, uppercase)", "[mixU:10]", 10)
	printExample("Alphanumerics (length = 10, mixed case)", "[mixR:10]", 10)
	printExample("Alphanumerics (length = 10, lowercase, short syntax)", "[m:10]", 10)
	printExample("Alphanumerics (length = 10, uppercase, short syntax)", "[mU:10]", 10)
	printExample("Alphanumerics (length = 10, mixed case, short syntax)", "[mR:10]", 10)
	printExample("List Of Integers", "[1..10]", 3)
	printExample("Strings From List", "[a,b,c,1,2,3]", 10)
	printExample("Strings From Random List", "[[str:1],[int:2],[mix:3]] - [#[[4..16],UUID]]", 10)
	printExample("Strings From Random List (short syntax)", "[[s:1],[i:2],[m:3]] - [#[[4..16],UUID]]", 10)
	printExample("Base64-Encoded Strings", "[b64:[str:20]]", 10)
	printExample("Base32-Encoded Strings", "[b32:[str:20]]", 10)
	printExample("ASCII85-Encoded Strings", "[a85:[str:20]]", 10)
	printExample("URL-Encoded Strings", "[url:[str:5]&[str:4]]", 10)
	printExample("Hex-Encoded Strings", "[hex:[str:5]&[str:4]]", 10)
	printExample("Binary-Encoded Strings", "[bin:[str:5]&[str:4]]", 10)
}

func main() {
	printExamples()
}
