package clubhouseapi

import "time"

// StartPhoneNumberAuthResponse is the response structure of the StartPhoneNumberAuth method
type StartPhoneNumberAuthResponse struct {
	Success      bool   `json:"success"`
	IsBlocked    bool   `json:"is_blocked"`
	ErrorMessage []byte `json:"error_message"`
}

// CompletePhoneNumberAuthResponse is the response structure of the CompletePhoneNumberAuth method
type CompletePhoneNumberAuthResponse struct {
	Success                   bool `json:"success"`
	IsVerified                bool `json:"is_verified"`
	NumberOfAttemptsRemaining int  `json:"number_of_attempts_remaining"`
	UserProfile               struct {
		UserID   int    `json:"user_id"`
		Name     string `json:"name"`
		PhotoURL string `json:"photo_url"`
		Username string `json:"username"`
	} `json:"user_profile"`
	AuthToken    string `json:"auth_token"`
	RefreshToken string `json:"refresh_token"`
	AccessToken  string `json:"access_token"`
	IsWaitlisted bool   `json:"is_waitlisted"`
	IsOnboarding bool   `json:"is_onboarding"`
}

type Channel struct {
	CreatorUserProfileID int    `json:"creator_user_profile_id"`
	ChannelID            int    `json:"channel_id"`
	Channel              string `json:"channel"`
	Topic                string `json:"topic"`
	IsPrivate            bool   `json:"is_private"`
	IsSocialMode         bool   `json:"is_social_mode"`
	URL                  string `json:"url"`
	Club                 struct {
		ClubID              int           `json:"club_id"`
		Name                string        `json:"name"`
		Description         string        `json:"description"`
		PhotoURL            string        `json:"photo_url"`
		NumMembers          int           `json:"num_members"`
		NumFollowers        int           `json:"num_followers"`
		EnablePrivate       bool          `json:"enable_private"`
		IsFollowAllowed     bool          `json:"is_follow_allowed"`
		IsMembershipPrivate bool          `json:"is_membership_private"`
		IsCommunity         bool          `json:"is_community"`
		Rules               []interface{} `json:"rules"`
		NumOnline           int           `json:"num_online"`
	} `json:"club"`
	ClubName              string      `json:"club_name"`
	ClubID                int         `json:"club_id"`
	WelcomeForUserProfile interface{} `json:"welcome_for_user_profile"`
	NumOther              int         `json:"num_other"`
	HasBlockedSpeakers    bool        `json:"has_blocked_speakers"`
	IsExploreChannel      bool        `json:"is_explore_channel"`
	NumSpeakers           int         `json:"num_speakers"`
	NumAll                int         `json:"num_all"`
	Users                 []struct {
		UserID              int       `json:"user_id"`
		Name                string    `json:"name"`
		PhotoURL            string    `json:"photo_url"`
		IsSpeaker           bool      `json:"is_speaker"`
		IsModerator         bool      `json:"is_moderator"`
		TimeJoinedAsSpeaker time.Time `json:"time_joined_as_speaker"`
		IsFollowedBySpeaker bool      `json:"is_followed_by_speaker"`
		IsInvitedAsSpeaker  bool      `json:"is_invited_as_speaker"`
	} `json:"users"`
}

type Event struct {
	EventID     int       `json:"event_id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	TimeStart   time.Time `json:"time_start"`
	Club        struct {
		ClubID              int           `json:"club_id"`
		Name                string        `json:"name"`
		Description         string        `json:"description"`
		PhotoURL            string        `json:"photo_url"`
		NumMembers          int           `json:"num_members"`
		NumFollowers        int           `json:"num_followers"`
		EnablePrivate       bool          `json:"enable_private"`
		IsFollowAllowed     bool          `json:"is_follow_allowed"`
		IsMembershipPrivate bool          `json:"is_membership_private"`
		IsCommunity         bool          `json:"is_community"`
		Rules               []interface{} `json:"rules"`
		NumOnline           int           `json:"num_online"`
	} `json:"club"`
	IsMemberOnly bool   `json:"is_member_only"`
	URL          string `json:"url"`
	Hosts        []struct {
		UserID   int    `json:"user_id"`
		Name     string `json:"name"`
		PhotoURL string `json:"photo_url"`
		Username string `json:"username"`
		Bio      string `json:"bio"`
		Twitter  string `json:"twitter"`
	} `json:"hosts"`
	Channel   interface{} `json:"channel"`
	IsExpired bool        `json:"is_expired"`
}

// GetChannelsResponse is the response structure of the GetChannels method
type GetChannelsResponse struct {
	Channels []Channel
	Events   []Event
	Success  bool `json:"success"`
}

// RefreshTokenResponse is the response structure of the RefreshToken method
type RefreshTokenResponse struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}

// JoinChannelResponse is the response structure of the JoinChannel method
type JoinChannelResponse struct {
	CreatorUserProfileID int    `json:"creator_user_profile_id"`
	ChannelID            int    `json:"channel_id"`
	Channel              string `json:"channel"`
	Topic                string `json:"topic"`
	IsPrivate            bool   `json:"is_private"`
	IsSocialMode         bool   `json:"is_social_mode"`
	URL                  string `json:"url"`
	Club                 struct {
		ClubID              int    `json:"club_id"`
		Name                string `json:"name"`
		Description         string `json:"description"`
		PhotoURL            string `json:"photo_url"`
		NumMembers          int    `json:"num_members"`
		NumFollowers        int    `json:"num_followers"`
		EnablePrivate       bool   `json:"enable_private"`
		IsFollowAllowed     bool   `json:"is_follow_allowed"`
		IsMembershipPrivate bool   `json:"is_membership_private"`
		IsCommunity         bool   `json:"is_community"`
		Rules               []struct {
			Desc  string `json:"desc"`
			Title string `json:"title"`
		} `json:"rules"`
		NumOnline int `json:"num_online"`
	} `json:"club"`
	ClubName              string      `json:"club_name"`
	ClubID                int         `json:"club_id"`
	WelcomeForUserProfile interface{} `json:"welcome_for_user_profile"`
	IsHandraiseEnabled    bool        `json:"is_handraise_enabled"`
	HandraisePermission   int         `json:"handraise_permission"`
	IsClubMember          bool        `json:"is_club_member"`
	IsClubAdmin           bool        `json:"is_club_admin"`
	Users                 []struct {
		UserID              int       `json:"user_id"`
		Name                string    `json:"name"`
		PhotoURL            string    `json:"photo_url"`
		Username            string    `json:"username"`
		FirstName           string    `json:"first_name"`
		Skintone            int       `json:"skintone"`
		IsNew               bool      `json:"is_new"`
		IsSpeaker           bool      `json:"is_speaker"`
		IsModerator         bool      `json:"is_moderator"`
		TimeJoinedAsSpeaker time.Time `json:"time_joined_as_speaker"`
		IsFollowedBySpeaker bool      `json:"is_followed_by_speaker"`
		IsInvitedAsSpeaker  bool      `json:"is_invited_as_speaker"`
	} `json:"users"`
	Success                 bool        `json:"success"`
	IsEmpty                 bool        `json:"is_empty"`
	Token                   string      `json:"token"`
	RtmToken                string      `json:"rtm_token"`
	PubnubToken             string      `json:"pubnub_token"`
	PubnubOrigin            interface{} `json:"pubnub_origin"`
	PubnubHeartbeatValue    int         `json:"pubnub_heartbeat_value"`
	PubnubHeartbeatInterval int         `json:"pubnub_heartbeat_interval"`
	PubnubEnable            bool        `json:"pubnub_enable"`
	AgoraNativeMute         bool        `json:"agora_native_mute"`
}

// LeaveChannelResponse is the response structure of the LeaveChannel method
type LeaveChannelResponse struct {
	Success bool `json:"success"`
}

// ActivePingResponse is the response structure of the ActivePing method
type ActivePingResponse struct {
	ShouldLeave bool `json:"should_leave"`
	Success     bool `json:"success"`
}
