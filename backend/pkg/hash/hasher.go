package hash

type Hasher interface {
	Hash(text string) (string, error)
	Check(text string, hash string) bool
}
