package expandurl

import "net/http"

//Expand URL
func Expand(uri string) (string, error) {
	resp, err := http.Get(uri)

	if err != nil {
		return "", err
	}

	return resp.Request.URL.String(), nil
}
