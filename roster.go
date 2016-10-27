package openfirerest

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type RosterGroup struct {
	Group string `json:"group" xml:"group"`
}

const (
	SubscriptionNone = "0"
	SubscriptionTo   = "1"
	SubscriptionFrom = "2"
	SubscriptionBoth = "3"
)

type RosterItem struct {
	JID              string      `json:"jid" xml:"jid"`
	NickName         string      `json:"nickname" xml:"nickname"`
	SubscriptionType string      `json:"subscriptionType" xml:"subscriptionType"`
	Groups           RosterGroup `json:"groups" xml:"groups"`
}

const EndpointPatternRoster string = "%s/plugins/restapi/v1/users/%s/roster"
const EndpointPatternRosterItem string = "%s/plugins/restapi/v1/users/%s/roster/%s"

func GetRoster(server, authorization, username string) ([]RosterItem, error) {
	url := fmt.Sprintf(EndpointPatternRoster, server, username)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	prepareRequest(req, authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	var items []RosterItem
	if resp.StatusCode != http.StatusOK {
		return items, errors.New("http response code isn't OK")
	}
	response, _ := ioutil.ReadAll(resp.Body)
	var objmap map[string]*json.RawMessage
	err = json.Unmarshal([]byte(response), &objmap)
	const rosterItemKey = "rosterItem"
	rawMsg, ok := objmap[rosterItemKey]
	if ok {
		if (*rawMsg)[0] == '[' {
			err = json.Unmarshal(*rawMsg, &items)
		} else {
			var item RosterItem
			err = json.Unmarshal(*rawMsg, &item)
			items = append(items, item)
		}
	}
	return items, err
}

func AddRoster(server, authorization, username string, roster RosterItem) error {
	url := fmt.Sprintf(EndpointPatternRoster, server, username)

	jsonStr, err := json.Marshal(roster)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	prepareRequest(req, authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode == http.StatusCreated {
		return nil
	}
	payload, _ := ioutil.ReadAll(resp.Body)
	return errors.New(string(payload))
}

func DeleteRoster(server, authorization, username, rosterJID string) error {
	url := fmt.Sprintf(EndpointPatternRosterItem, server, username, rosterJID)

	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return err
	}
	prepareRequest(req, authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	payload, _ := ioutil.ReadAll(resp.Body)
	return errors.New(string(payload))
}

func UpdateRoster(server, authorization, username string, roster RosterItem) error {
	url := fmt.Sprintf(EndpointPatternRosterItem, server, username, roster.JID)

	jsonStr, err := json.Marshal(roster)
	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}
	prepareRequest(req, authorization)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode == http.StatusOK {
		return nil
	}
	payload, _ := ioutil.ReadAll(resp.Body)
	return errors.New(string(payload))
}
