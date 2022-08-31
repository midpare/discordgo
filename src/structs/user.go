package structs

import (
	. "discordbot/src/utils"
)

type UserFlags uint
type UserPremiumType uint

const (
	UserFlags_Staff UserFlags = 1 << iota
	UserFlags_Partner
	UserFlags_HypeSquad
	UserFlags_BugHunterLevel1
	UserFlags_HypesquadOnlineHouse1
	UserFlags_HypesquadOnlineHouse2
	UserFlags_HypesquadOnlineHouse3
	UserFlags_PremiumEarlySupporter
	UserFlags_TeamPseudoUser
	UserFlags_BugHunterLevel2
	UserFlags_VerifiedBot
	UserFlags_VerifiedDeveloper
	UserFlags_CertifiedModerator
	UserFlags_BotHttpInteractions
)

const (
	None UserPremiumType = iota
	NitroClassic
	Nitro
)

type User struct {
	Id            Snowflake   `json:"id"`
	Username      Snowflake   `json:"username"`
	Discriminator string      `json:"discriminator"`
	Avatar        string      `json:"avatar"`
	Bot           bool        `json:"bot"`
	System        bool        `json:"system"`
	MfaEnabled    bool        `json:"mfa_enabled"`
	Banner        string      `json:"banner"`
	AccentColor   string      `json:"accent_color"`
	Locale        string      `json:"locale"`
	Verified      bool        `json:"verified"`
	Email         string      `json:"email"`
	Flags         UserFlags   `json:"flags"`
	PremiumType   PremiumTier `json:"premium_type"`
	PublicFlags   UserFlags   `json:"public_flags"`
}
