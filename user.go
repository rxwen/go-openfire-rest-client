package openfirerest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	if resp.StatusCode != http.StatusOK {
		return false, nil
	}
	response, _ := ioutil.ReadAll(resp.Body)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal([]byte(response), &objmap)
	const nameKey = "username"
	rawUsername, ok := objmap[nameKey]
	if !ok {
		return false, nil
	} else {
		var returnedUsername string
		err = json.Unmarshal(*rawUsername, &returnedUsername)
		if err != nil || returnedUsername != username {
			return false, err
		}
	}
	return true, nil
}
