package encryptor

type Encryptor interface {
	Encrypt(text string) (string, error)
	Decrypt(encryptedText string) (string, error)
}
