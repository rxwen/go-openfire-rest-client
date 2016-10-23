package openfirerest

type RosterItem struct {
	JID              string   `json:"jid" xml:"jid"`
	NickName         string   `json:"nickname" xml:"nickname"`
	SubscriptionType int      `json:"subscription_type" xml:"subscriptionType"`
	Groups           []string `json:"groups" xml:"groups"`
}

const EntpointPatternRoster string = "plugins/restapi/v1/users/%s/roster"
