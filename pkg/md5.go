package pkg

import (
	"crypto/md5"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func GetMD5(filePath string) string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	h := md5.New()
	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}
	md5Str := fmt.Sprintf("%x", h.Sum(nil))

	return strings.ToUpper(md5Str)
}
