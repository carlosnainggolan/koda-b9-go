package service

import (
	"fmt"
	"io"
	"os"
)

func ReadingFile(path string) error {
	open, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("File path tidak sesuai")
	}

	defer open.Close()

	ctn, err := io.ReadAll(open)
	if err != nil {
		return  fmt.Errorf("Tidak bisa membaca direktori (harus file)")
	}

	println(string(ctn))
	return nil
}