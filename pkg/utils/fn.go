package utils

import (
	"math/rand"
	"net"
	"net/http"
	"strings"
	"time"
)

func RandomBytes(lengthParam ...int) []byte {
	length := 16
	if len(lengthParam) > 0 {
		length = lengthParam[0]
	}
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	var res []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		res = append(res, bytes[r.Intn(len(bytes))])
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
