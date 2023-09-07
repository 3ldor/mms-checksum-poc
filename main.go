package main

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"strings"

	"golang.org/x/text/encoding/unicode"
)

func main() {
	token := "Don'tMessWithMMS"
	payload := ""   // base64 encoded payload
	signature := "" // base64 encoded signature

	// Build plaintext
	plaintext := payload[10:20] + token + signature[2:10]

	// Convert to UTF-16
	encoder := unicode.UTF16(unicode.LittleEndian, unicode.IgnoreBOM).NewEncoder()
	data, _ := encoder.String(plaintext)

	// Hash using SHA1
	hash := sha1.Sum([]byte(data))

	// Select 8 specific bytes and hex encode
	checksum := strings.ToUpper(hex.EncodeToString(hash[2:10]))

	fmt.Println(checksum)
}
