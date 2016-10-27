package openfirerest

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const EndpointPatternProperty string = "%s/plugins/restapi/v1/system/properties/%s"

func GetProperty(server, authorization, name string) (string, error) {
	url := fmt.Sprintf(EndpointPatternProperty, server, name)

	var value string
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return value, err
	}
	prepareRequest(req, authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return value, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode == http.StatusNotFound {
		return value, errors.New(fmt.Sprintf("property %s not found", name))
	}
	if resp.StatusCode != http.StatusOK {
		return value, errors.New("http response code isn't OK")
	}
	response, _ := ioutil.ReadAll(resp.Body)
	var objmap map[string]string
	err = json.Unmarshal([]byte(response), &objmap)
	if err != nil {
		return value, err
	}
	const keyOfValue = "@value"
	value, ok := objmap[keyOfValue]
	if !ok {
		return value, errors.New("invalid response")
	}
	return value, err
}
