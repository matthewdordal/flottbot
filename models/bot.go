package models

import "github.com/sirupsen/logrus"

// Bot is a struct representation of bot.yml
type Bot struct {
	// Bot fields
	ID      string                            `mapstructure:"id"`
	Name    string                            `mapstructure:"name" binding:"required"`
	Remotes map[string]map[string]interface{} `mapstructure:"remotes" binding:"required"`
	// SlackToken                    string            `mapstructure:"slack_token"`
	// SlackVerificationToken        string            `mapstructure:"slack_verification_token"`
	// SlackWorkspaceToken           string            `mapstructure:"slack_workspace_token"`
	// SlackEventsCallbackPath       string            `mapstructure:"slack_events_callback_path"`
	// SlackInteractionsCallbackPath string            `mapstructure:"slack_interactions_callback_path"`
	// DiscordToken                  string            `mapstructure:"discord_token"`
	// Users                         map[string]string `mapstructure:"slack_users"`
	// UserGroups                    map[string]string `mapstructure:"slack_usergroups"`
	// Channels                      map[string]string `mapstructure:"slack_channels"`
	// CLI                   bool   `mapstructure:"cli,omitempty"`
	// CLIUser               string `mapstructure:"cli_user,omitempty"`
	// Scheduler             bool   `mapstructure:"scheduler,omitempty"`
	Debug   bool `mapstructure:"debug,omitempty"`
	LogJSON bool `mapstructure:"log_json,omitempty"`
	// InteractiveComponents bool   `mapstructure:"interactive_components,omitempty"`
	Metrics        bool   `mapstructure:"metrics,omitempty"`
	CustomHelpText string `mapstructure:"custom_help_text,omitempty"`
	// System
	Log          logrus.Logger
	RunChat      bool
	RunCLI       bool
	RunScheduler bool
}
