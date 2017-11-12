package protos

type user struct {
	Pk                         int64  `json:"pk"`
	Username                   string `json:"username"`
	FullName                   string `json:"full_name"`
	IsPrivate                  bool   `json:"is_private"`
	ProfilePicURL              string `json:"profile_pic_url"`
	IsVerified                 bool   `json:"is_verified"`
	HasAnonymousProfilePicture bool   `json:"has_anonymous_profile_picture"`
}

type loggedInUser struct {
	user
	IsBusiness            bool   `json:"is_business"`
	CanSeeOrganicInsights bool   `json:"can_see_organic_insights"`
	ShowInsightsTerms     bool   `json:"show_insights_terms"`
	AllowContactsSync     bool   `json:"allow_contacts_sync"`
	PhoneNumber           string `json:"phone_number"`
}

type threadUser struct {
	user
	FriendshipStatus friendshipStatus `json:"friendship_status"`
}

type friendshipStatus struct {
	Following       bool `json:"following"`
	Blocking        bool `json:"blocking"`
	IsPrivate       bool `json:"is_private"`
	IncomingRequest bool `json:"incoming_request"`
	OutgoingRequest bool `json:"outgoing_request"`
	IsBestie        bool `json:"is_bestie"`
}
