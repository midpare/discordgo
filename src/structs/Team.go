package structs

import . "discordbot/src/utils"

type MembershipStateType uint

const (
	MembershipStateType_Invited MembershipStateType = iota
	MembershipStateType_Accepted
)

type Team struct {
	Icon        string        `json:"icon"`
	Id          Snowflake     `json:"id"`
	Members     []*TeamMember `json:"members"`
	Name        string        `json:"name"`
	OwnerUserId Snowflake     `json:"owner_user_id"`
}

type TeamMember struct {
	MembershipState MembershipStateType `json:"membership_state"`
	Permissions     []string            `json:"permissions"`
	TeamId          Snowflake           `json:"team_id"`
	User            *User               `json:"user"`
}
