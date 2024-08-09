package cdnjs

import (
	"crypto/tls"
	"encoding/json"
	"io"
	"net/http"
	"net/url"
)

// https://api.cdnjs.com/libraries/{:library}
// https://api.cdnjs.com/libraries/jquery

// https://api.cdnjs.com/libraries/{:library}/{:version}
// https://api.cdnjs.com/libraries/jquery/3.5.1

// https://cdnjs.cloudflare.com/ajax/libs/{:library}/{:version}/{:file}

func Get[T any](fullUrl, proxy string, ret *T) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}
		tr.Proxy = http.ProxyURL(proxyURL)
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(fullUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respData, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(respData, &ret)
	if err != nil {
		return nil, err
	}

	return respData, nil
}

func GetFileBytes(fullUrl, proxy string) ([]byte, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	if proxy != "" {
		proxyURL, err := url.Parse(proxy)
		if err != nil {
			return nil, err
		}
		tr.Proxy = http.ProxyURL(proxyURL)
	}

	client := &http.Client{Transport: tr}

	resp, err := client.Get(fullUrl)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
