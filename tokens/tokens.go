package tokens

import (
	"bufio"
	"log"
	"os"
	"strings"

	"github.com/toxyl/go-string-generator/utils"
)

const (
	tokenPatternHex               string = "abcdef0123456789"
	tokenPatternInt               string = "0123456789"
	tokenPatternStringLower       string = "abcdefghijklmnopqrstuvwxyz"
	tokenPatternStringUpper       string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tokenPatternStringMix         string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	tokenPatternAlphanumericLower string = "abcdefghijklmnopqrstuvwxyz0123456789"
	tokenPatternAlphanumericUpper string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	tokenPatternAlphanumericMix   string = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

type Tokens struct {
	tokens []string
}

func (ts *Tokens) Length() int {
	return len(ts.tokens)
}

func (ts *Tokens) Random() string {
	return ts.tokens[utils.GetRandomInt(0, len(ts.tokens)-1)]
}

func (ts *Tokens) FromFile(file string) *Tokens {
	if file == "" {
		return ts
	}

	readFile, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	readFile.Close()
	for _, line := range lines {
		line = strings.Trim(line, "\t\r\n ")
		if line == "" || line[0] == '#' {
			continue
		}
		ts.tokens = append(ts.tokens, line)
	}
	return ts
}

func (ts *Tokens) FromSlice(tokens []string) *Tokens {
	for _, token := range tokens {
		token = strings.Trim(token, "\t\r\n ")
		if token == "" || token[0] == '#' {
			continue
		}
		ts.tokens = append(ts.tokens, token)
	}
	return ts
}

func (ts *Tokens) Append(tokens ...string) *Tokens {
	for _, token := range tokens {
		token = strings.Trim(token, "\t\r\n ")
		if token == "" || token[0] == '#' {
			continue
		}
		ts.tokens = append(ts.tokens, token)
	}
	return ts
}
