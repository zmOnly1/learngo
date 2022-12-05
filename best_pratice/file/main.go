package main

import (
	"fmt"
	"io/fs"
	"path/filepath"
)

func ListDirectory(dir string) ([]string, error) {
	return nil, nil
}

func ListDirectory1(dir string) chan string {
	return nil
}

func main() {
	filepath.Walk(".", func(path string, info fs.FileInfo, err error) error {
		fmt.Printf("path: %s\n", path)
		return nil
	})
}
