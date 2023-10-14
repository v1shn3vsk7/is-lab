package hash

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"github.com/htruong/go-md2"
	"github.com/rs/zerolog/log"
	"io"
)

func EncryptSecret(secret string, salt []byte) (string, string) {
	secretB := []byte(secret)

	block, err := aes.NewCipher(SecretKey)
	if err != nil {
		log.Err(err)
		return "", ""
	}

	ciphertext := make([]byte, len(secretB))

	if len(salt) == 0 {
		salt = make([]byte, aes.BlockSize)
		_, err = io.ReadFull(rand.Reader, salt)
		if err != nil {
			log.Err(err)
			return "", ""
		}
	}

	saltStr := base64.StdEncoding.EncodeToString(salt)

	cfb := cipher.NewCFBEncrypter(block, salt)
	cfb.XORKeyStream(ciphertext, secretB)

	md2Hash := md2.New()
	md2Hash.Write(SecretKey)
	keyHash := md2Hash.Sum(nil)

	ciphertext = append(keyHash, ciphertext...)

	return hex.EncodeToString(ciphertext), saltStr
}
