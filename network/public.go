package network

import (
	"io/ioutil"
	"net/http"
)

func getPublicIp() (string, error) {
	resp, err := http.Get("https://api.ipify.org/?format=text")

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	ip, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return "", err
	}

	return string(ip), nil
}
