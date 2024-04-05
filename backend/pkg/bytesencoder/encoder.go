package bytesencoder

type Encoder interface {
	Encode(b []byte) string
	Decode(text string) ([]byte, error)
}
