package tokens

import (
	"fmt"

	"github.com/toxyl/go-string-generator/utils"
)

type TokenIntRange struct {
	Min int
	Max int
}

func (t *TokenIntRange) Parse() string {
	return fmt.Sprintf("%d", utils.GetRandomInt(t.Min, t.Max))
}
