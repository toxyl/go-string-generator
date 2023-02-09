package tokens

import "github.com/toxyl/go-string-generator/utils"

type TokenMixUpper struct {
	Length int
}

func (t *TokenMixUpper) Parse() string {
	return utils.GetRandomString(tokenPatternAlphanumericUpper, t.Length)
}
