package model

type Post struct {
	Pk                   int64           `json:"pk"`
	ID                   string          `json:"id"`
	TextPostAppInfo      TextPostAppInfo `json:"text_post_app_info"`
	Caption              Caption         `json:"caption"`
	TakenAt              int             `json:"taken_at"`
	DeviceTimestamp      int             `json:"device_timestamp"`
	MediaType            int             `json:"media_type"`
	Code                 string          `json:"code"`
	ClientCacheKey       string          `json:"client_cache_key"`
	FilterType           int             `json:"filter_type"`
	ProductType          string          `json:"product_type"`
	OrganicTrackingToken string          `json:"organic_tracking_token"`
	ImageVersions2       struct {
		Candidates []struct {
			Width        int    `json:"width"`
			Height       int    `json:"height"`
			URL          string `json:"url"`
			ScansProfile string `json:"scans_profile"`
		} `json:"candidates"`
	} `json:"image_versions2"`
	OriginalWidth             int      `json:"original_width"`
	OriginalHeight            int      `json:"original_height"`
	VideoVersions             []string `json:"video_versions"`
	LikeCount                 int      `json:"like_count"`
	TimezoneOffset            int      `json:"timezone_offset"`
	HasLiked                  bool     `json:"has_liked"`
	LikeAndViewCountsDisabled bool     `json:"like_and_view_counts_disabled"`
	CanViewerReshare          bool     `json:"can_viewer_reshare"`
	IntegrityReviewDecision   string   `json:"integrity_review_decision"`
	TopLikers                 []string `json:"top_likers"`
	User                      User     `json:"user"`
}

type TextPostAppInfo struct {
	IsPostUnavailable     bool        `json:"is_post_unavailable"`
	IsReply               bool        `json:"is_reply"`
	ReplyToAuthor         interface{} `json:"reply_to_author"`
	DirectReplyCount      int         `json:"direct_reply_count"`
	SelfThreadCount       int         `json:"self_thread_count"`
	ReplyFacepileUsers    []User      `json:"reply_facepile_users"`
	LinkPreviewAttachment interface{} `json:"link_preview_attachment"`
	CanReply              bool        `json:"can_reply"`
	ReplyControl          string      `json:"reply_control"`
	HushInfo              interface{} `json:"hush_info"`
	ShareInfo             ShareInfo   `json:"share_info"`
}

type ShareInfo struct {
	CanRepost          bool `json:"can_repost"`
	IsRepostedByViewer bool `json:"is_reposted_by_viewer"`
	CanQuotePost       bool `json:"can_quote_post"`
}

type User struct {
	HasAnonymousProfilePicture     bool        `json:"has_anonymous_profile_picture"`
	FanClubInfo                    FanClubInfo `json:"fan_club_info"`
	FBIDV2                         interface{} `json:"fbid_v2"`
	TransparencyProductEnabled     bool        `json:"transparency_product_enabled"`
	TextPostAppTakeABreakSetting   int         `json:"text_post_app_take_a_break_setting"`
	InteropMessagingUserFBID       int64       `json:"interop_messaging_user_fbid"`
	ShowInsightsTerms              bool        `json:"show_insights_terms"`
	AllowedCommenterType           string      `json:"allowed_commenter_type"`
	IsUnpublished                  bool        `json:"is_unpublished"`
	ReelAutoArchive                string      `json:"reel_auto_archive"`
	CanBoostPost                   bool        `json:"can_boost_post"`
	CanSeeOrganicInsights          bool        `json:"can_see_organic_insights"`
	HasOnboardedToTextPostApp      bool        `json:"has_onboarded_to_text_post_app"`
	TextPostAppJoinerNumber        int         `json:"text_post_app_joiner_number"`
	Pk                             int64       `json:"pk"`
	PKID                           string      `json:"pk_id"`
	Username                       string      `json:"username"`
	FullName                       string      `json:"full_name"`
	IsPrivate                      bool        `json:"is_private"`
	ProfilePicURL                  string      `json:"profile_pic_url"`
	AccountBadges                  []string    `json:"account_badges"`
	FeedPostReshareDisabled        bool        `json:"feed_post_reshare_disabled"`
	ShowAccountTransparencyDetails bool        `json:"show_account_transparency_details"`
	ThirdPartyDownloadsEnabled     int         `json:"third_party_downloads_enabled"`
}

type Caption struct {
	Pk                 string `json:"pk"`
	UserID             int64  `json:"user_id"`
	Text               string `json:"text"`
	Type               int    `json:"type"`
	CreatedAt          int64  `json:"created_at"`
	CreatedAtUTC       int64  `json:"created_at_utc"`
	ContentType        string `json:"content_type"`
	Status             string `json:"status"`
	BitFlags           int    `json:"bit_flags"`
	DidReportAsSpam    bool   `json:"did_report_as_spam"`
	ShareEnabled       bool   `json:"share_enabled"`
	User               User   `json:"user"`
	IsCovered          bool   `json:"is_covered"`
	IsRankedComment    bool   `json:"is_ranked_comment"`
	MediaID            int64  `json:"media_id"`
	PrivateReplyStatus int    `json:"private_reply_status"`
}

type FanClubInfo struct {
	FanClubID                            int64       `json:"fan_club_id"`
	FanClubName                          string      `json:"fan_club_name"`
	IsFanClubReferralEligible            interface{} `json:"is_fan_club_referral_eligible"`
	FanConsiderationPageRevampEligiblity interface{} `json:"fan_consideration_page_revamp_eligiblity"`
	IsFanClubGiftingEligible             interface{} `json:"is_fan_club_gifting_eligible"`
	SubscriberCount                      interface{} `json:"subscriber_count"`
	ConnectedMemberCount                 interface{} `json:"connected_member_count"`
	AutosaveToExclusiveHighlight         interface{} `json:"autosave_to_exclusive_highlight"`
	HasEnoughSubscribersForSSC           interface{} `json:"has_enough_subscribers_for_ssc"`
}
