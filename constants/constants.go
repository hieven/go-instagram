package constants

import (
	"bytes"
	"html/template"
	"reflect"
)

const (
	SIG_KEY     = "68a04945eb02970e2e8d15266fc256f7295da123e123f44b88f09d594a5902df"
	SIG_VERSION = "4"
	APP_VERSION = "10.8.0"
)

var (
	HOSTNAME     = "i.instagram.com"
	WEB_HOSTNAME = "www.instagram.com"
	HOST         = "https://" + HOSTNAME + "/"
	WEBHOST      = "https://" + WEB_HOSTNAME + "/"
	API_ENDPOINT = HOST + "api/v1/"
)

func GetURL(name string, data interface{}) string {
	t := template.New("url template")

	r := reflect.ValueOf(ROUTES)
	f := reflect.Indirect(r).FieldByName(name).String()
	t, _ = t.Parse(f)

	var url bytes.Buffer
	t.Execute(&url, data)

	return url.String()
}

var ROUTES = struct {
	HOSTNAME     string
	WEB_HOSTNAME string
	HOST         string
	WEBHOST      string

	ThreadsBroadcastText string
	Inbox                string
	Login                string
	LocationFeed         string
	ThreadsApproveAll    string
	ThreadsShow          string
	TimelineFeed         string
	Like                 string
	Unlike               string
}{
	HOSTNAME:     HOSTNAME,
	WEB_HOSTNAME: WEB_HOSTNAME,
	HOST:         HOST,
	WEBHOST:      WEBHOST,

	ThreadsBroadcastText: API_ENDPOINT + "direct_v2/threads/broadcast/text/",
	Inbox:                API_ENDPOINT + "direct_v2/inbox/{{if .Cursor}}?cursor={{.Cursor}}{{end}}",
	Login:                API_ENDPOINT + "accounts/login/",
	LocationFeed:         API_ENDPOINT + "feed/location/",
	ThreadsApproveAll:    API_ENDPOINT + "direct_v2/threads/approve_all/",
	ThreadsShow:          API_ENDPOINT + "direct_v2/threads/",
	TimelineFeed:         API_ENDPOINT + "feed/timeline/?rank_token={{.RankToken}}{{if .MaxID}}&max_id={{.MaxID}}{{end}}&ranked_content=true",
	Like:                 API_ENDPOINT + "media/{{.ID}}/like/",
	Unlike:               API_ENDPOINT + "media/{{.ID}}/unlike/",
}
