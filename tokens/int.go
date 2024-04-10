package tokens

import "github.com/toxyl/go-string-generator/utils"

type TokenInt struct {
	Length int
}

func (t *TokenInt) Parse() string {
	return utils.GetRandomString(tokenPatternInt, t.Length)
}
