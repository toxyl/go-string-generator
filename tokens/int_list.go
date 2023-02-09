package tokens

import (
	"fmt"
	"strings"
)

type TokenIntList struct {
	Min int
	Max int
}

func (t *TokenIntList) Parse() string {
	a := make([]string, t.Max-t.Min+1)
	for i := range a {
		a[i] = fmt.Sprint(t.Min + i)
	}
	return strings.Join(a, ",")
}
