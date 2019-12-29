package constants

import "fmt"

const (
	SigCsrfToken = "missing"
	SigDeviceID  = "android-b256317fd493b848"
	SigKey       = "c36436a942ea1dbb40d7f2d7d45280a620d991ce8c62fb4ce600f0a048c32c11"
	SigVersion   = "4"
	AppVersion   = "107.0.0.27.121"

	Scheme   = "https"
	Hostname = "i.instagram.com"

	InstagramStatusFail = "fail"
)

var (
	Host        = fmt.Sprintf("%s://%s", Scheme, Hostname)
	APIEndpoint = fmt.Sprintf("%s/api/v1", Host)

	LoginEndpoint        = APIEndpoint + "/accounts/login/"
	InboxEndpoint        = APIEndpoint + "/direct_v2/inbox/"
	PendingInboxEndpoint = APIEndpoint + "/direct_v2/pending_inbox/"

	TimelineFeedEndpoint = APIEndpoint + "/feed/timeline/?ranked_content=true"

	ThreadApproveAllEndpoint      = APIEndpoint + "/direct_v2/threads/approve_all/"
	ThreadApproveMultipleEndpoint = APIEndpoint + "/direct_v2/threads/approve_multiple/"
	ThreadBroadcastTextEndpoint   = APIEndpoint + "/direct_v2/threads/broadcast/text/"
	ThreadBroadcastLinkEndpoint   = APIEndpoint + "/direct_v2/threads/broadcast/link/"
	ThreadBroadcastShareEndpoint  = APIEndpoint + "/direct_v2/threads/broadcast/media_share/?media_type=photo"
	ThreadShowEndpoint            = APIEndpoint + "/direct_v2/threads/%s/"

	MediaInfoEndpoint   = APIEndpoint + "/media/%s/info/"
	MediaLikeEndpoint   = APIEndpoint + "/media/%s/like/"
	MediaUnlikeEndpoint = APIEndpoint + "/media/%s/unlike/"

	LocationFeedEndpoint    = APIEndpoint + "/feed/location/%d/"
	LocationSectionEndpoint = APIEndpoint + "/locations/%d/sections/"
)
