package models

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/hieven/go-instagram/constants"
	"github.com/hieven/go-instagram/utils"
)

type TimelineFeed struct {
	Items         []*FeedItem `json:"feed_items"`
	MoreAvailable bool        `json:"-"`
	Cursor        string      `json:"-"`
	Instagram     *Instagram  `json:"-"`
}

type FeedItem struct {
	MediaOrAd      MediaOrAd      `json:"media_or_ad"`
	SuggestedUsers SuggestedUsers `json:"suggested_users"`
}

type MediaOrAd struct {
	Pk              int64           `json:"pk"`
	ID              string          `json:"id"`
	MediaType       int             `json:"media_type"`
	ImageVersions2  ImageVersions2  `json:"image_versions2"`
	Caption         Caption         `json:"caption"`
	CaptionIsEdited bool            `json:"caption_is_edited"`
	VideoVersions   []*VideoVersion `json:"video_versions"`
	VideoDuriation  float64         `json:"video_duration"`
	User            User            `json:"user"`
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

type timelineFeedResponse struct {
	MoreAvailable bool   `json:"more_available"`
	NextMaxID     string `json:"next_max_id"`
	TimelineFeed
	DefaultResponse
}

func (feed *TimelineFeed) SetCursor(maxID string) {
	feed.Cursor = maxID
}

func (feed *TimelineFeed) GetCursor() string {
	return feed.Cursor
}

func (feed *TimelineFeed) isMoreAvailable() bool {
	return feed.MoreAvailable
}

func (feed *TimelineFeed) Get() ([]*FeedItem, error) {
	userID := strconv.FormatInt(feed.Instagram.Pk, 10)

	rankToken := utils.GenerateRankToken(userID)

	url := constants.ROUTES.TimelineFeed + "&rank_token=" + rankToken + "&max_id=" + feed.Cursor

	agent := feed.Instagram.AgentPool.Get()
	defer feed.Instagram.AgentPool.Put(agent)

	_, body, _ := feed.Instagram.SendRequest(agent.Get(url))

	var resp timelineFeedResponse
	json.Unmarshal([]byte(body), &resp)

	if resp.Status == "fail" {
		return nil, errors.New(resp.Message)
	}

	feed.Items = resp.Items
	feed.MoreAvailable = resp.MoreAvailable

	if feed.MoreAvailable {
		feed.Cursor = resp.NextMaxID
	}

	return feed.Items, nil
}
