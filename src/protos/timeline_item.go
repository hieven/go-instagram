package protos

type TimelineItem struct {
	MediaOrAd      MediaOrAd      `json:"media_or_ad"`
	SuggestedUsers SuggestedUsers `json:"suggested_users"`
}

type MediaOrAd struct {
	Pk              int64           `json:"pk"`
	ID              string          `json:"id"`
	MediaType       int             `json:"media_type"`
	ImageVersions2  *ImageVersions2 `json:"image_versions2"`
	Caption         *Caption        `json:"caption"`
	CaptionIsEdited bool            `json:"caption_is_edited"`
	VideoVersions   []*VideoVersion `json:"video_versions"`
	VideoDuriation  float64         `json:"video_duration"`
	User            MediaOrAdUser   `json:"user"`
	HasMoreComments bool            `json:"has_more_comments"`
	HasLiked        bool            `json:"has_liked"`
	HasAudio        bool            `json:"has_audio"`
	NextMaxID       int64           `json:"next_max_id"`
	ViewCount       int             `json:"view_count"`
	CommentCount    int             `json:"comment_count"`
	LikeCount       int             `json:"like_count"`
}

type SuggestedUsers struct {
	Title            string `json:"title"`
	ViewAllText      string `json:"view_all_text"`
	LandingSiteTitle string `json:"landing_site_title"`
	LandingSiteType  string `json:"landing_site_type"`
	Type             int    `json:"type"`
	TrackingToken    string `json:"tracking_token"`
}
