package expandurl

import "net/http"
import "net/url"

//Expand URL
func Expand(uri string) (string, error) {

	decodedURL, urlError := url.QueryUnescape(uri)

	if urlError != nil {
		return "", urlError
	}

	resp, err := http.Get(decodedURL)

	if err != nil {
		return "", err
	}

	return resp.Request.URL.String(), nil
}
