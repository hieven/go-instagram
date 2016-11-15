package models

type Media struct {
	Pk              int            `json:"pk"`
	ID              string         `json:"id"`
	DeviceTimestamp int            `json:"device_timestamp"`
	MediaType       int            `json:"media_type"`
	Code            string         `json:"code"`
	ImageVersions2  ImageVersions2 `json:"image_versions2"`
	Location        Location       `json:"location"`
	LikeCount       int            `json:"like_count"`
	HasLiked        bool           `json:"has_liked"`
}
