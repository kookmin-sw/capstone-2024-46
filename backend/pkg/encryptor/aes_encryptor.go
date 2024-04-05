package encryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"errors"
	"io"

	"private-llm-backend/pkg/bytesencoder"
	"private-llm-backend/pkg/errorutil"
)

var _ Encryptor = (*aesEncryptor)(nil)

type aesEncryptor struct {
	aes     cipher.Block
	encoder bytesencoder.Encoder
}

// Encrypt is a method that encrypts the given text using the AES block cipher algorithm in CBC mode with PKCS#7.
// The encrypted data is then returned as a base64-encoded string.
//
// Parameters:
// - text: The text to be encrypted.
//
// Returns:
// - string: The encrypted text as a base64-encoded string.
// - error: An error if encryption fails.
func (v *aesEncryptor) Encrypt(text string) (string, error) {
	// Make sure the text is padded to the block size
	paddedText := []byte(text)
	blockSize := v.aes.BlockSize()

	// Add paddings
	padding := blockSize - len(paddedText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	paddedText = append(paddedText, padText...)

	// Generate a random initialization vector (IV)
	iv := make([]byte, blockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", errorutil.WithDetail(err, errors.New("failed to create iv"))
	}

	// cipherText will hold the encrypted data
	cipherText := make([]byte, len(paddedText))

	// Create a new Cipher Block Chain (CBC) mode
	mode := cipher.NewCBCEncrypter(v.aes, iv)
	mode.CryptBlocks(cipherText, paddedText)

	result := append(iv, cipherText...)

	encodedResult := v.encoder.Encode(result)

	return encodedResult, nil
}

// Decrypt is a method that decrypts a given base64-encoded string using the AES block cipher algorithm in CBC mode with PKCS#7.
// The decrypted data is then returned as a string.
//
// Parameters:
// - encryptedText: The base64-encoded string to be decrypted.
//
// Returns:
// - string: The decrypted text as a string.
// - error: An error if decryption fails.
func (v *aesEncryptor) Decrypt(encryptedText string) (string, error) {
	inputBytes, err := v.encoder.Decode(encryptedText)
	if err != nil {
		return "", errorutil.WithDetail(err, errors.New("failed to decode encrypted text"))
	}

	blockSize := v.aes.BlockSize()
	if len(inputBytes) < blockSize*2 {
		return "", errorutil.Error(errors.New("encrypted text is too short"))
	}

	// IV is the first blockSize bytes of the inputBytes
	iv := inputBytes[:blockSize]
	cipherText := inputBytes[blockSize:]

	// plainText will hold the decrypted data
	plainText := make([]byte, len(cipherText))
	mode := cipher.NewCBCDecrypter(v.aes, iv)
	mode.CryptBlocks(plainText, cipherText)

	// padding 제거
	padding := plainText[len(plainText)-1]
	plainText = plainText[:len(plainText)-int(padding)]

	return string(plainText), nil
}

func NewAESEncryptor(secretInBase64 string, encoder bytesencoder.Encoder) (Encryptor, error) {
	secret, err := base64.StdEncoding.DecodeString(secretInBase64)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to decode secret"))
	}
	key := make([]byte, 16)
	copy(key, secret)
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, errorutil.WithDetail(err, errors.New("failed to create aes cipher"))
	}
	return &aesEncryptor{
		aes:     c,
		encoder: encoder,
	}, nil
}
