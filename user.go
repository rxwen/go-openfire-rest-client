package openfirerest

import (
	"fmt"
	"net/http"
)

const EndpointPatternUser string = "%s/plugins/restapi/v1/users/%s"

func IsUserExist(server, authorization, username string) (bool, error) {
	url := fmt.Sprintf(EndpointPatternUser, server, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return false, err
	}
	prepareRequest(req, authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode == http.StatusNotFound {
		return false, nil
	} else {
		return true, nil
	}
}
