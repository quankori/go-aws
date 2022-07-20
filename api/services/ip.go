package services

import (
	"io/ioutil"
	"net/http"

	"github.com/quankori/go-aws/configs"
)

// LocalIP get the host machine local IP address
func LocalIP() string {
	config, _ := configs.LoadConfig()
	resp, err := http.Get(config.IpURI)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	ip, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	msg := string(ip[:])
	return msg
}
