package gostringgenerator

import "github.com/toxyl/go-string-generator/tokens"

func NewGenerator(dataDir string, errorHandler func(err error)) *tokens.RandomStringGenerator {
	return tokens.NewRandomStringGenerator(dataDir, errorHandler)
}

func NewGeneratorSimple() *tokens.RandomStringGenerator {
	return tokens.NewRandomStringGeneratorSimple()
}
