package packets

import . "discordbot/src/utils"

type VerificationLevel uint
type DefaultMessageNotificationLevel uint
type ExplicitContentFilterLevel uint
type GuildFeature string
type MFALevel uint
type SystemChannelFlags uint
type PremiumTier uint
type GuildNSFWLevel uint

const (
	VerificationLevel_None VerificationLevel = iota
	VerificationLevel_Low
	VerificationLevel_Medium
	VerificationLevel_High
	VerificationLevel_VeryHigh
)

const (
	DefaultMessageNotificationLevel_AllMembers DefaultMessageNotificationLevel = iota
	DefaultMessageNotificationLevel_OnlyMentions
)
const (
	ExplicitContentFilterLevel_Disabled ExplicitContentFilterLevel = iota
	ExplicitContentFilterLevel_MembersWithoutRoles
	ExplicitContentFilterLevel_AllMembers
)

const (
	GuildFeature_AnimatedBanner               = GuildFeature("guild has access to set an animated guild banner image")
	GuildFeature_AnimatedIcon                 = GuildFeature("guild has access to set an animated guild icon")
	GuildFeature_AutoModeration               = GuildFeature("guild has set up auto moderation rules")
	GuildFeature_Banner                       = GuildFeature("guild has access to set a guild banner image")
	GuildFeature_Community                    = GuildFeature("guild can enable welcome screen, Membership Screening, stage channels and discovery, and receives community updates")
	GuildFeature_Discoverable                 = GuildFeature("guild is able to be discovered in the directory")
	GuildFeature_Featurable                   = GuildFeature("guild is able to be featured in the directory")
	GuildFeature_InviteSplash                 = GuildFeature("guild has access to set an invite splash background")
	GuildFeature_MemberVerificationGateEnbled = GuildFeature("guild has enabled Membership Screening")
	GuildFeature_MonetizationEnabled          = GuildFeature("guild has enabled monetization")
	GuildFeature_MoreStickers                 = GuildFeature("guild has increased custom sticker slots")
	GuildFeature_News                         = GuildFeature("guild has access to create news channels")
	GuildFeature_Partnered                    = GuildFeature("guild is partnered")
	GuildFeature_PreviewEnabled               = GuildFeature("guild can be previewed before joining via Membership Screening or the directory")
	GuildFeature_PrivateThreads               = GuildFeature("guild has access to create private threads")
	GuildFeature_RoleIcons                    = GuildFeature("guild is able to set role icons")
	GuildFeature_TicketedEventsEnabled        = GuildFeature("guild has enabled ticketed events")
	GuildFeature_VanityUrl                    = GuildFeature("guild has access to set a vanity URL")
	GuildFeature_Verified                     = GuildFeature("guild is verified")
	GuildFeature_VipRegions                   = GuildFeature("guild has access to set 384kbps bitrate in voice")
	GuildFeature_WelcomeScreenEnabled         = GuildFeature("guild has enabled the welcome screen")
)

const (
	MFALevel_None MFALevel = iota
	MFALevel_EleVated
)

const (
	SystemChannelFlags_SuppressJoinNotifications SystemChannelFlags = 1 << iota
	SystemChannelFlags_SuppressPremiumSubscriptions
	SystemChannelFlags_SuppressGuildReminderNotifications
	SystemChannelFlags_SuppressJoinNotificationReplies
)

const (
	PremiumTier_None PremiumTier = iota
	PremiumTier_Tier1
	PremiumTier_Tier2
	PremiumTier_Tier3
)

const (
	GuildNSFWLevel_Default GuildNSFWLevel = iota
	GuildNSFWLevel_Explicit
	GuildNSFWLevel_Safe
	GuildNSFWLevel_AgeRestricted
)

type Guild struct {
	Id                          Snowflake                       `json:"id"`
	Name                        string                          `json:"name"`
	Icon                        string                          `json:"icon"`
	IconHash                    string                          `json:"icon_hash"`
	Splash                      string                          `json:"splash"`
	DiscoverySplash             string                          `json:"discovery_splash"`
	Owner                       bool                            `json:"owner"`
	OwnerId                     string                          `json:"owner_id"`
	Permissions                 string                          `json:"permissions"`
	Region                      string                          `json:"region"`
	AfkChannelId                Snowflake                       `json:"afk_channel_id"`
	AfkTimeout                  int                             `json:"afk_timeout"`
	WidgetEnabled               bool                            `json:"widget_enabled"`
	WidgetChannelId             Snowflake                       `json:"widget_channel_id"`
	VerificationLevel           VerificationLevel               `json:"verification_level"`
	DefaultMessageNotifications DefaultMessageNotificationLevel `json:"default_message_notifications"`
	ExplicitContentFilter       ExplicitContentFilterLevel      `json:"explicit_content_filter"`
	Roles                       []*Role                         `json:"roles"`
	Emojis                      []*Emoji                        `json:"emojis"`
	Features                    []GuildFeature                  `json:"features"`
	MFALevel                    *MFALevel                       `json:"mfa_level"`
	ApplicationId               Snowflake                       `json:"application_id"`
	SystemChannelId             Snowflake                       `json:"system_channel_id"`
	SystemChannelFlags          SystemChannelFlags              `json:"system_channel_flags"`
	RulesChannelId              Snowflake                       `json:"rules_channel_id"`
	MaxPresences                int                             `json:"max_presences"`
	MaxMembers                  int                             `json:"max_members"`
	VanityUrlCode               string                          `json:"vanity_url_code"`
	Description                 string                          `json:"description"`
	Banner                      string                          `json:"banner"`
	PremiumTier                 PremiumTier                     `json:"premium_tier"`
	PremiumSubscriptionCount    int                             `json:"premium_subscription_count"`
	PreferredLocale             string                          `json:"preferred_locale"`
	PublicUpdatesChannelId      Snowflake                       `json:"public_updates_channel_id"`
	MaxVideoChannelUsers        int                             `json:"max_video_channel_id"`
	ApproximateMemberCount      int                             `json:"approximate_member_count"`
	ApproximatePresenceCount    int                             `json:"approximate_presence_count"`
	WelcomeScreen               *WelcomeScreen                  `json:"welcome_screen"`
	NSFWLevel                   GuildNSFWLevel                  `json:"nsfw_level"`
	Stickers                    []*Sticker                      `json:"sticker"`
	PremiumProgressBarEnabled   bool                            `json:"premium_progress_bar_enabled"`
	JoinedAt                    string                          `json:"joined_at"`
	Large                       bool                            `json:"large"`
	Unavailable                 bool                            `json:"unavailable"`
	MemberCount                 int                             `json:"member_count"`
	VoiceStates                 []*VoiceState                   `json:"voice_states"`
	Members                     []*GuildMember                  `json:"members"`
	Channels                    []*Channel                      `json:"channels"`
	Threads                     []*Channel                      `json:"threads"`
	Presences                   []*PresencesUpdate              `json:"presences"`
	StageInstances              []*StageInstance                `json:"stage_instances"`
	GuildScheduledEvents        []*GuildScheduledEvent          `json:"guild_scheduled_events"`
}

type WelcomeScreen struct {
	Description     string                  `json:"description"`
	WelcomeChannels []*WelcomeScreenChannel `json:"welcome_channels"`
}

type WelcomeScreenChannel struct {
	ChannelId   Snowflake `json:"channel_id"`
	Description string    `json:"description"`
	EmojiId     Snowflake `json:"emoji_id"`
	EmojiName   string    `json:"emoji_name"`
}
