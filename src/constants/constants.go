package constants

import "fmt"

const (
	SigCsrfToken = "missing"
	SigDeviceID  = "android-b256317fd493b848"
	SigKey       = "0443b39a54b05f064a4917a3d1da4d6524a3fb0878eacabf1424515051674daa"
	SigVersion   = "4"
	AppVersion   = "10.33.0"

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
