package openfirerest

// requires precense plugin to work: https://www.igniterealtime.org/projects/openfire/plugins/presence/readme.html
import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const (
	PrecenseOnline  = "online"
	PrecenseOffline = "offline"
)

const EndpointPatternPrecense string = "%s/plugins/presence/status"

func GetPrecense(server, authorization, jid string) (string, error) {
	url := fmt.Sprintf(EndpointPatternPrecense, server) //, username)

	req, err := http.NewRequest("GET", url, nil)
	precense := PrecenseOffline
	if err != nil {
		return precense, err
	}
	q := req.URL.Query()
	q.Add("jid", jid)
	q.Add("type", "text")
	req.URL.RawQuery = q.Encode()

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return precense, err
	}
	defer func() {
		if resp != nil {
			resp.Body.Close()
		}
	}()

	if resp.StatusCode != http.StatusOK {
		return precense, errors.New("http response code isn't OK")
	}
	response, _ := ioutil.ReadAll(resp.Body)
	r := string(response)
	if len(r) > 0 {
		r = r[0 : len(r)-1]
	}
	if -1 == strings.Index(r, "Unavailable") {
		precense = PrecenseOnline
	}
	return precense, err
}
