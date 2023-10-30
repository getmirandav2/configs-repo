package config

import (
	"path"
	"runtime"
	"sync"
)

var (
	rootDir string
	once    sync.Once
)

func getRootDir() {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return
	}
	rootDir = path.Join(path.Dir(filename), "..", "..")
}

// RootDir returns the root directory of the project.
func RootDir() string {
	once.Do(func() {
		getRootDir()
	})
	return rootDir
}
