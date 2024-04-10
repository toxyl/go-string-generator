package tokens

import "github.com/toxyl/go-string-generator/utils"

type TokenStr struct {
	Length int
}

func (t *TokenStr) Parse() string {
	return utils.GetRandomString(tokenPatternStringMix, t.Length)
}
