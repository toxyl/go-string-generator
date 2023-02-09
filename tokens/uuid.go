package tokens

import (
	"github.com/toxyl/go-string-generator/utils"
)

type TokenRandomUUID string

func (vir *TokenRandomUUID) Parse() string {
	return utils.GetRandomString(tokenPatternHex, 8) +
		"-" + utils.GetRandomString(tokenPatternHex, 4) +
		"-" + utils.GetRandomString(tokenPatternHex, 4) +
		"-" + utils.GetRandomString(tokenPatternHex, 12)
}
