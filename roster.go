package openfirerest

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RosterGroup struct {
	Group string `json:"group" xml:"group"`
}

type RosterItem struct {
	JID              string        `json:"jid" xml:"jid"`
	NickName         string        `json:"nickname" xml:"nickname"`
	SubscriptionType int           `json:"subscriptionType" xml:"subscriptionType"`
	Groups           []RosterGroup `json:"groups" xml:"groups"`
}

const EndpointPatternRoster string = "plugins/restapi/v1/users/%s/roster"

func GetRoster(username string) ([]RosterItem, error) {
	url := "http://127.0.0.1:9090/" + fmt.Sprintf(EndpointPatternRoster, username)

	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	response, _ := ioutil.ReadAll(resp.Body)
	var rosters []RosterItem
	err = json.Unmarshal(response, rosters)
	return rosters, err
}

func AddRoster(username string, roster RosterItem) error {
	url := "http://127.0.0.1:9090/" + fmt.Sprintf(EndpointPatternRoster, username)

	fmt.Println(url)
	jsonStr, err := json.Marshal(roster)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Authorization", "Basic YWRtaW46YWRtaW4=")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	payload, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(payload))
	return nil
}
