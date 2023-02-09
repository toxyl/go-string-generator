package tokens

import "github.com/toxyl/go-string-generator/utils"

type TokenMixLower struct {
	Length int
}

func (t *TokenMixLower) Parse() string {
	return utils.GetRandomString(tokenPatternAlphanumericLower, t.Length)
}
