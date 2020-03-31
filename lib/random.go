package lib

import (
	"github.com/brianvoe/gofakeit"
	"math/rand"
	"time"
)

var alpha = "abcdefghijkmnpqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ23456789"

//RandomString generate random strings
func RandomString(size int) string {

	rand.Seed(time.Now().UTC().UnixNano())

	buf := make([]byte, size)
	for i := 0; i < size; i++ {
		buf[i] = alpha[rand.Intn(len(alpha))]
	}
	return string(buf)
}

//RandomPort generate random strings
func RandomPort() int {
	port := []int{25, 443, 80}
	return port[rand.Intn(3)]
}

//RandomProtocol generate random strings
func RandomProtocol() string {
	protocol := []string{"SMTP", "WEB"}
	return protocol[rand.Intn(2)]
}

//RandomRequest generate random strings
func RandomRequest() string {
	request := []string{"<n/a>", "GET", "POST"}
	return request[rand.Intn(3)]
}

//RandomURL generate random URL strings
func RandomURL() string {
	url := []string{"<n/a>", gofakeit.URL()}
	return url[rand.Intn(2)]
}
