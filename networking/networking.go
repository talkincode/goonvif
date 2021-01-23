package networking

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func SendSoap(endpoint, message string) (*http.Response, error) {
	if os.Getenv("GOVNIF_DEBUG" ) == "true" || os.Getenv("GOVNIF_DEBUG" ) == "1" {
		fmt.Fprintln(os.Stdout, message)
	}
	httpClient := new(http.Client)

	resp, err := httpClient.Post(endpoint, "application/soap+xml; charset=utf-8", bytes.NewBufferString(message))
	if err != nil {
		return resp, err
	}

	return resp, nil
}
