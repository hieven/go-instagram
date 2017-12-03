package protos

type Media struct {
	Pk                   int             `json:"pk"`
	ID                   string          `json:"id"`
	MediaType            int             `json:"media_type"`
	FilterType           int             `json:"filter_type"`
	CarouselMedia        []*Media        `json:"carousel_media"`
	ImageVersions2       *ImageVersions2 `json:"image_versions2"`
	Location             *Location       `json:"location"`
	OriginalWidth        int             `json:"original_width"`
	OriginalHeight       int             `json:"original_height"`
	Lat                  float64         `json:"lat"`
	Lng                  float64         `json:"lng"`
	Code                 string          `json:"code"`
	LikeCount            int             `json:"like_count"`
	CommentCount         int             `json:"comment_count"`
	Caption              *Caption        `json:"caption"`
	HasLiked             bool            `json:"has_liked"`
	HasMoreComments      bool            `json:"has_more_comments"`
	ClientCacheKey       string          `json:"client_cache_key"`
	OrganicTrackingToken string          `json:"organic_tracking_token"`
	DeviceTimestamp      int             `json:"device_timestamp"`
	TakenAt              int             `json:"taken_at"`
}
