package constants

const (
	SIG_KEY     = "fc4720e1bf9d79463f62608c86fbddd374cc71bbfb98216b52e3f75333bd130d"
	SIG_VERSION = "4"
	APP_VERSION = "9.4.0"
)

var HOSTNAME = "i.instagram.com"
var WEB_HOSTNAME = "www.instagram.com"
var HOST = "https://" + HOSTNAME + "/"
var WEBHOST = "https://" + WEB_HOSTNAME + "/"
var API_ENDPOINT = HOST + "api/v1/"

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
}{
	HOSTNAME:     HOSTNAME,
	WEB_HOSTNAME: WEB_HOSTNAME,
	HOST:         HOST,
	WEBHOST:      WEBHOST,

	ThreadsBroadcastText: API_ENDPOINT + "direct_v2/threads/broadcast/text/",
	Inbox:                API_ENDPOINT + "direct_v2/inbox/",
	Login:                API_ENDPOINT + "accounts/login/",
	LocationFeed:         API_ENDPOINT + "feed/location/",
	ThreadsApproveAll:    API_ENDPOINT + "direct_v2/threads/approve_all/",
	ThreadsShow:          API_ENDPOINT + "direct_v2/threads/",
	TimelineFeed:         API_ENDPOINT + "feed/timeline/?ranked_content=true",
}
