package utils

import (
	"math"
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

func RandomStr(lengthParam ...int) string {
	length := 16
	if len(lengthParam) > 0 {
		length = lengthParam[0]
	}
	bytes := []byte{
		'0', '1', '2', '3', '4', '5', '6', '7', '8', '9',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm',
		'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	var res []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	sourceLen := len(bytes)
	for i := 0; i < length; i++ {
		res = append(res, bytes[r.Intn(sourceLen)])
	}
	return string(res)
}
func RandomBytes(lengthParam ...int) []byte {
	length := 16
	if len(lengthParam) > 0 {
		length = lengthParam[0]
	}
	var res []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		res = append(res, byte(r.Intn(math.MaxUint8)))
	}
	return res
}
func GetClientIP(req *http.Request) string {
	clientIP := req.Header.Get("X-Forwarded-For")
	clientIP = strings.TrimSpace(strings.Split(clientIP, ",")[0])
	if clientIP == "" {
		clientIP = strings.TrimSpace(req.Header.Get("X-Real-Ip"))
	}
	if clientIP != "" {
		return clientIP
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(req.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
