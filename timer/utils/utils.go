package utils

import (
	"io/ioutil"
	"net/http"
)

func GetHtmlLength(url string) (uint, error) {
	response, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	return (uint)(len(body)), nil
}
