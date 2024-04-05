package hash

import "golang.org/x/crypto/bcrypt"

var _ Hasher = (*passwordHasher)(nil)

type passwordHasher struct {
}

func (v *passwordHasher) Hash(t string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(t), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (v *passwordHasher) Check(t string, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(t))
	return err == nil
}

func NewPasswordHasher() Hasher {
	return &passwordHasher{}
}
