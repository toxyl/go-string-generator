package tokens

import "github.com/toxyl/go-string-generator/utils"

type TokenHash struct {
	Length int
}

func (t *TokenHash) Parse() string {
	return utils.GetRandomString(tokenPatternHex, t.Length)
}
