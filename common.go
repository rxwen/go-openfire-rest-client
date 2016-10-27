package openfirerest

import "net/http"

func prepareRequest(req *http.Request, authorization string) {
	req.Header.Set("Authorization", authorization)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
}
