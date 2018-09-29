package http

import (
	"io/ioutil"
	"net/http"
)

func HttpGet(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()
	res, err := ioutil.ReadAll(resp.Body)
	return res, err
}
