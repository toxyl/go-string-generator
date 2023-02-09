package tokens

import "github.com/toxyl/go-string-generator/utils"

type TokenStrUpper struct {
	Length int
}

func (t *TokenStrUpper) Parse() string {
	return utils.GetRandomString(tokenPatternStringUpper, t.Length)
}
