package constants

var HOSTNAME = "i.instagram.com"
var WEB_HOSTNAME = "www.instagram.com"
var HOST = "https://" + HOSTNAME + "/"
var WEBHOST = "https://" + WEB_HOSTNAME + "/"

var ThreadsBroadcastText = HOST + "api/v1/direct_v2/threads/broadcast/text/"
var Inbox = HOST + "api/v1/direct_v2/inbox/"
var Login = HOST + "api/v1/accounts/login/"

var ROUTES = struct {
	HOSTNAME     string
	WEB_HOSTNAME string
	HOST         string
	WEBHOST      string

	ThreadsBroadcastText string
	Inbox                string
	Login                string
}{
	HOSTNAME:     HOSTNAME,
	WEB_HOSTNAME: WEB_HOSTNAME,
	HOST:         HOST,
	WEBHOST:      WEBHOST,

	ThreadsBroadcastText: ThreadsBroadcastText,
	Inbox:                Inbox,
	Login:                Login,
}
