package constants

import "fmt"

const (
	SigCsrfToken = "missing"
	SigDeviceID  = "android-b256317fd493b848"
	SigKey       = "68a04945eb02970e2e8d15266fc256f7295da123e123f44b88f09d594a5902df"
	SigVersion   = "4"
	AppVersion   = "10.8.0"

	Scheme   = "https"
	Hostname = "i.instagram.com"

	InstagramStatusFail = "fail"
)

var (
	Host        = fmt.Sprintf("%s://%s", Scheme, Hostname)
	APIEndpoint = fmt.Sprintf("%s/api/v1", Host)

	LoginEndpoint = fmt.Sprintf("%s/accounts/login/", APIEndpoint)
	InboxEndpoint = fmt.Sprintf("%s/direct_v2/inbox/", APIEndpoint)

	ThreadBroadcastTextEndpoint = fmt.Sprintf("%s/direct_v2/threads/broadcast/text/", APIEndpoint)
	ThreadBroadcastLinkEndpoint = fmt.Sprintf("%s/direct_v2/threads/broadcast/link/", APIEndpoint)
)
