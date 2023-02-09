package tokens

import "github.com/toxyl/go-string-generator/utils"

type TokenStrLower struct {
	Length int
}

func (t *TokenStrLower) Parse() string {
	return utils.GetRandomString(tokenPatternStringLower, t.Length)
}
