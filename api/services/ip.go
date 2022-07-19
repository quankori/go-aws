package services

import (
	"io/ioutil"
	"net/http"
)

// LocalIP get the host machine local IP address
func LocalIP() []byte {
	url := "https://api.ipify.org?format=text"
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	return ip
}
