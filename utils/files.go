package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// DirExists reports whether the dir exists as a boolean,
// taken from https://stackoverflow.com/a/49697453 / https://stackoverflow.com/a/51870143/3337885
func DirExists(name string) bool {
	fileOrDir, err := os.Open(name)
	if err != nil {
		return false
	}
	info, err := fileOrDir.Stat()
	if err != nil {
		return false
	}
	if info.IsDir() {
		return true
	}
	return false
}

func FileExists(name string) bool {
	file, err := os.Open(name)
	if err != nil {
		return false
	}
	_, err = file.Stat()
	return err == nil
}

func CopyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, in)
	if err != nil {
		return err
	}
	return out.Close()
}

func ReadFile(filename string) ([]byte, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func ListFiles(directory string) []string {
	files := []string{}
	fs, err := ioutil.ReadDir(directory)
	if err == nil {
		for _, file := range fs {
			if file.IsDir() {
				files = append(files, ListFiles(fmt.Sprintf("%s/%s", directory, file.Name()))...)
				continue
			}
			files = append(files, fmt.Sprintf("%s/%s", directory, file.Name()))
		}
	}
	return files
}
