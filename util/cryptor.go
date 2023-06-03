package util

import (
	"github.com/ProtonMail/go-crypto/openpgp"
	"io"
	"log"
	"os"
	"strings"
)

func Decrypt(fileName string, privateKey string) string {
	encryptFile, err := os.Open(fileName)
	if err != nil {
		defer encryptFile.Close()
	}

	privateKey = strings.ReplaceAll(privateKey, "\\n", "\n")
	entityList, err := openpgp.ReadArmoredKeyRing(strings.NewReader(privateKey))
	if err != nil {
		log.Fatal("read privateKey error", err)
	}

	md, err := openpgp.ReadMessage(encryptFile, entityList, nil, nil)
	if err != nil {
		log.Fatal("decrypt fail", err)
	}
	plaintext, err := io.ReadAll(md.UnverifiedBody)
	if err != nil {
		log.Fatal(err)
	}
	return string(plaintext)
}
