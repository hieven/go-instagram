package constants

import "fmt"

const (
	SigCsrfToken = "missing"
	SigDeviceID  = "android-b256317fd493b848"
	SigKey       = "109513c04303341a7daf27bb41b268e633b30dcc65a3fe14503f743176113869"
	SigVersion   = "4"
	AppVersion   = "27.0.0.7.97"

	Scheme   = "https"
	Hostname = "i.instagram.com"

	InstagramStatusFail = "fail"
)

var (
	Host        = fmt.Sprintf("%s://%s", Scheme, Hostname)
	APIEndpoint = fmt.Sprintf("%s/api/v1", Host)

	LoginEndpoint = APIEndpoint + "/accounts/login/"
	InboxEndpoint = APIEndpoint + "/direct_v2/inbox/"

	TimelineFeedEndpoint = APIEndpoint + "/feed/timeline/?ranked_content=true"

	ThreadApproveAllEndpoint     = APIEndpoint + "/direct_v2/threads/approve_all/"
	ThreadBroadcastTextEndpoint  = APIEndpoint + "/direct_v2/threads/broadcast/text/"
	ThreadBroadcastLinkEndpoint  = APIEndpoint + "/direct_v2/threads/broadcast/link/"
	ThreadBroadcastShareEndpoint = APIEndpoint + "/direct_v2/threads/broadcast/media_share/?media_type=photo"
	ThreadShowEndpoint           = APIEndpoint + "/direct_v2/threads/%s/"

	MediaInfoEndpoint   = APIEndpoint + "/media/%s/info/"
	MediaLikeEndpoint   = APIEndpoint + "/media/%s/like/"
	MediaUnlikeEndpoint = APIEndpoint + "/media/%s/unlike/"

	LocationFeedEndpoint = APIEndpoint + "/feed/location/%d/"
)
