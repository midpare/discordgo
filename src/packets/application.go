package packets

import . "discordbot/src/utils"

type Application struct {
	Id                  Snowflake      `json:"id"`
	Name                string         `json:"name"`
	Icon                string         `json:"icon"`
	Desciption          string         `json:"description"`
	RpcOrigins          []string       `json:"rpc_origins"`
	BotPublic           bool           `json:"bot_public"`
	BotRequireCodeGrant bool           `json:"bot_require_code_grant"`
	TermsOfServiceUrl   string         `json:"terms_of_service_rul"`
	PrivacyPolicyUrl    string         `json:"privacy_policy_rul"`
	Owner               *User          `json:"owner"`
	VerifyKey           string         `json:"verify_key"`
	Team                *Team          `json:"team"`
	GuildId             Snowflake      `json:"guild_id"`
	PrimarySkuId        Snowflake      `json:"primary_sku_id"`
	Slug                string         `json:"slug"`
	CoverImage          string         `json:"cover_image"`
	Flags               int            `json:"flags"`
	Tags                []string       `json:"tags"`
	InstallParams       *InstallParams `json:"install_params"`
	CustomInstallUrl    string         `json:"custom_install_url"`
}

type InstallParams struct {
	Scopes      []string `json:"scopes"`
	Permissions string   `json:"permissions"`
}
