package tokens

import (
	"fmt"

	"github.com/toxyl/go-string-generator/utils"
)

type TokenRandomUUID string

func (vir *TokenRandomUUID) Parse() string {
	return fmt.Sprintf(
		"%s-%s-%s-%s",
		utils.GetRandomString(tokenPatternHex, 8),
		utils.GetRandomString(tokenPatternHex, 4),
		utils.GetRandomString(tokenPatternHex, 4),
		utils.GetRandomString(tokenPatternHex, 12),
	)
}
