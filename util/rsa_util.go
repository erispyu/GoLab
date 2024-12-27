package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
)

// RsaSignWithPrivateKeyPem sign by private key pem(should be PKSC1v1.5 format), with SHA256, the result is base64 encoded
func RsaSignWithPrivateKeyPem(data string, prvKeyPem []byte) (string, error) {
	if len(prvKeyPem) == 0 {
		return "", errors.New("prvKeyPem cannot be empty")
	}
	prvKey, _ := pem.Decode(prvKeyPem)
	var parsedKey interface{}
	parsedKey, err := x509.ParsePKCS1PrivateKey(prvKey.Bytes)
	if err != nil {
		parsedKey, err = x509.ParsePKCS8PrivateKey(prvKey.Bytes)
		if err != nil {
			return "", err
		}
	}
	privateKey, ok := parsedKey.(*rsa.PrivateKey)
	if !ok {
		return "", fmt.Errorf("not *rsa.PrivateKey error")
	}
	h := sha256.New()
	h.Write([]byte(data))
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, hash)
	if err != nil {
		return "", err
	}
	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}
