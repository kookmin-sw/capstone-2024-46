package bytesencoder

import "encoding/base64"

var _ Encoder = (*urlSafeEncoder)(nil)

type urlSafeEncoder struct {
	encoder *base64.Encoding
}

const customBase64Chars = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789-_"

func (v *urlSafeEncoder) Encode(b []byte) string {
	return v.encoder.EncodeToString(b)
}

func (v *urlSafeEncoder) Decode(text string) ([]byte, error) {
	return v.encoder.DecodeString(text)
}

func NewUrlSafeEncoder() Encoder {
	return &urlSafeEncoder{
		encoder: base64.NewEncoding(customBase64Chars).WithPadding(base64.NoPadding),
	}
}
