package main

import (
	"fmt"
	"os"
)
import "path/filepath"

var SUBMODULES = []string{
	"http",
	"utils",
	"kubernetes",
}

func main() {
	for _, module := range SUBMODULES {
		filepath.Walk(filepath.Join(module), func(path string, info os.FileInfo, err error) error {
			fmt.Println(path, info, err)
			return nil
		})
	}
}
