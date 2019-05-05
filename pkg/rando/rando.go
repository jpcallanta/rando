package rando

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func GetFiles(dir string) []string {
	var files []string

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		log.Fatal(err)
	}

	return files
}

func RenameFiles(files []string, dry bool) {
	var randoFileName string

	for _, f := range files {
		timeHash := strconv.FormatInt(time.Now().UnixNano(), 10)

		if strings.HasPrefix(filepath.Base(f), ".") {
			randoFileName = filepath.Dir(f) + "/" + "." + CreateRandomFileName(f, timeHash)
		} else {
			randoFileName = filepath.Dir(f) + "/" + CreateRandomFileName(f, timeHash) + filepath.Ext(f)
		}

		if dry {
			fmt.Println(f, "->", randoFileName)
		} else {
			err := os.Rename(f, randoFileName)

			fmt.Println(f, "->", randoFileName)

			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func CreateRandomFileName(file string, t string) string {
	hasher := md5.New()

	hasher.Write([]byte(t + file))

	return hex.EncodeToString(hasher.Sum(nil))
}
