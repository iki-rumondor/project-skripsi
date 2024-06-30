package utils

import (
	"fmt"
	"math/rand"
	"mime/multipart"
	"strings"
	"time"
)

func GenerateRandomString(length int) string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomFileName(file *multipart.FileHeader) string {

	fileExt := strings.ToLower(file.Filename[strings.LastIndex(file.Filename, ".")+1:])
	randString := GenerateRandomString(12)

	return fmt.Sprintf("%s.%s", randString, fileExt)
}
