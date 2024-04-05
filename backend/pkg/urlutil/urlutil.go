package urlutil

import "net/url"

func NewURLStringWithParams(u url.URL, path string, params map[string]string) string {
	q := make(url.Values)
	for k, v := range params {
		q.Set(k, v)
	}
	u.Path = path
	u.RawQuery = q.Encode()
	return u.String()
}
