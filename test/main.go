package main

import (
	"fmt"
	"os"
)

func main() {
	generateFilename("jpg")
}

func generateFilename(ext string) string {
	f, err := os.CreateTemp(".", "*."+ext)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	fmt.Println(f.Name())
	return f.Name()
}
