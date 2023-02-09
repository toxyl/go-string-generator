package utils

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/radovskyb/watcher"
)

type FilesCache struct {
	lock         *sync.Mutex
	files        map[string][]string
	dirlist      map[string][]string
	errorHandler func(err error)
}

func (fc *FilesCache) watch(path string) {
	w := watcher.New()
	w.IgnoreHiddenFiles(true)
	w.FilterOps(watcher.Write, watcher.Create)

	go func() {
		for {
			select {
			case event := <-w.Event:
				if !event.IsDir() {
					fc.getFile(event.Path, true)
				}
			case err := <-w.Error:
				fc.errorHandler(err)
			case <-w.Closed:
				return
			}
		}
	}()

	if err := w.AddRecursive(path); err != nil {
		log.Fatalln(err)
	}

	if err := w.Start(time.Millisecond * 1000); err != nil {
		log.Fatalln(err)
	}
}

func (fc *FilesCache) hasFile(file string) bool {
	fc.lock.Lock()
	defer fc.lock.Unlock()
	if _, ok := fc.files[file]; ok {
		return true
	}
	return false
}

func (fc *FilesCache) hasDirlist(dir string) bool {
	fc.lock.Lock()
	defer fc.lock.Unlock()
	if _, ok := fc.dirlist[dir]; ok {
		return true
	}
	return false
}

func (fc *FilesCache) getDirlist(dir string) []string {
	fc.lock.Lock()
	defer fc.lock.Unlock()
	if _, ok := fc.dirlist[dir]; ok {
		return fc.dirlist[dir]
	}
	return nil
}

func (fc *FilesCache) setDirlist(dir string, files []string) {
	fc.lock.Lock()
	defer fc.lock.Unlock()
	fc.dirlist[dir] = files
}

func (fc *FilesCache) getFile(file string, forceLoad bool) []string {
	exists := fc.hasFile(file)

	if forceLoad || !exists {
		f, err := os.Open(file)
		if err != nil {
			fc.errorHandler(fmt.Errorf("failed to open '%s'", file))
			return []string{}
		}
		defer f.Close()

		fc.lock.Lock()
		fc.files[file] = []string{}
		fc.lock.Unlock()

		scanner := bufio.NewScanner(f)

		for scanner.Scan() {
			t := strings.Trim(scanner.Text(), " \t\r\n")

			if t != "" && t[0] != '#' { // must not be empty and line must not start with #
				fc.lock.Lock()
				fc.files[file] = append(fc.files[file], t)
				fc.lock.Unlock()
			}
		}
	}

	fc.lock.Lock()
	defer fc.lock.Unlock()
	return fc.files[file]
}

func (fc *FilesCache) GetRandomLineFromFile(file string) string {
	if !fc.hasFile(file) {
		if !FileExists(file) {
			return ""
		}

		if DirExists(file) {
			if !fc.hasDirlist(file) {
				fc.setDirlist(file, ListFiles(file))
			}
			file = GetRandomStringFromList(fc.getDirlist(file)...)
		}
	}

	data := fc.getFile(file, false)

	l := len(data) - 1
	if l < 0 {
		return ""
	}

	return data[GetRandomInt(0, l)]
}

func NewFilesCache(path string, errorHandler func(err error)) *FilesCache {
	fc := &FilesCache{
		lock:         &sync.Mutex{},
		files:        map[string][]string{},
		dirlist:      map[string][]string{},
		errorHandler: errorHandler,
	}
	go func() {
		fc.watch(path)
	}()

	return fc
}
