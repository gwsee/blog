package common

var Jwt jwt

type jwt struct{}

func (jwt) BuildToken(data string, key string, expire int64) string {
	return ""
}
