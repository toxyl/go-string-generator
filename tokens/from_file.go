package tokens

import (
	"path/filepath"

	"github.com/toxyl/go-string-generator/utils"
)

var FilesCache *utils.FilesCache

type TokenLineFromFile struct {
	File string
}

func (t *TokenLineFromFile) Parse(dataDir string) string {
	path, err := filepath.Rel(dataDir, filepath.Join(dataDir, t.File))
	if err != nil {
		return "" // silently ignore
	}
	if path == "" || (len(path) >= 3 && path[0:3] == "../") {
		return ""
	}

	return FilesCache.GetRandomLineFromFile(filepath.Join(dataDir, path))
}
