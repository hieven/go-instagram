package constants

var HOSTNAME = "i.instagram.com"
var WEB_HOSTNAME = "www.instagram.com"
var HOST = "https://" + HOSTNAME + "/"
var WEBHOST = "https://" + WEB_HOSTNAME + "/"
var API_ENDPOINT = HOST + "api/v1/"

var ThreadsBroadcastText = API_ENDPOINT + "direct_v2/threads/broadcast/text/"
var Inbox = API_ENDPOINT + "direct_v2/inbox/"
var Login = API_ENDPOINT + "accounts/login/"
var LocationFeed = API_ENDPOINT + "feed/location/"
var ThreadsApproveAll = API_ENDPOINT + "direct_v2/threads/approve_all/"
var ThreadsShow = API_ENDPOINT + "direct_v2/threads/"

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
}{
	HOSTNAME:     HOSTNAME,
	WEB_HOSTNAME: WEB_HOSTNAME,
	HOST:         HOST,
	WEBHOST:      WEBHOST,

	ThreadsBroadcastText: ThreadsBroadcastText,
	Inbox:                Inbox,
	Login:                Login,
	LocationFeed:         LocationFeed,
	ThreadsApproveAll:    ThreadsApproveAll,
	ThreadsShow:          ThreadsShow,
}
