package models

// Action defines the structure for Actions used within Rules
type Action struct {
	Name             string                 `mapstructure:"name" binding:"required"`
	Type             string                 `mapstructure:"type" binding:"required"`
	URL              string                 `mapstructure:"url"`
	Cmd              string                 `mapstructure:"cmd"`
	Timeout          int                    `mapstructure:"timeout"`
	QueryData        map[string]interface{} `mapstructure:"query_data"`
	CustomHeaders    map[string]string      `mapstructure:"custom_headers"`
	Auth             []Auth                 `mapstructure:"auth"`
	ExposeJSONFields map[string]string      `mapstructure:"expose_json_fields"`
	Response         string                 `mapstructure:"response"`
	LimitToChannels  []string               `mapstructure:"limit_to_channels"`
	Message          string                 `mapstructure:"message"`
	Reaction         string                 `mapstructure:"update_reaction" binding:"omitempty"`
}

// Auth is a basic Auth data structure
type Auth struct {
	Type string `mapstructure:"type"`
	User string `mapstructure:"user"`
	Pass string `mapstructure:"pass"`
}
