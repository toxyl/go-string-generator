package tokens

import "github.com/toxyl/go-string-generator/utils"

type TokenMix struct {
	Length int
}

func (t *TokenMix) Parse() string {
	return utils.GetRandomString(tokenPatternAlphanumericMix, t.Length)
}
