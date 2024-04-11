package tokens

import (
	"github.com/toxyl/go-string-generator/utils"
)

type TokenStrFromList struct {
	Strings []string
}

func (t *TokenStrFromList) Parse() string {
	return utils.GetRandomStringFromList(t.Strings...)
}
