package lib

import (
	"crypto/tls"
	"io/ioutil"
	"net/http"
	"time"
)

// RequestURL 请求指定Url
func RequestURL(url string) (*http.Response, error) {
	// to configure TLS, need use http.Transport and http.Client
	tr := &http.Transport{
		// connection timeout
		ResponseHeaderTimeout: time.Second * 3,
		// disable tls verify
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true,
		},
	}
	client := &http.Client{
		Transport: tr,
	}

	response, err := client.Get(url)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	return response, err
}

func getRespBody(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()

	// read http response body
	body, err := ioutil.ReadAll(resp.Body)
	return body, err
}
