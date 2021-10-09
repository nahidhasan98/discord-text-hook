package discord_text_hook

import "time"

type User struct {
	ID            string `json:"id"`
	Username      string `json:"username"`
	Discriminator string `json:"discriminator"`
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
}

type Message struct {
	ID              string    `json:"id"`
	ChannelID       string    `json:"channel_id"`
	GuildID         string    `json:"guild_id"`
	Author          User      `json:"author"`
	Content         string    `json:"content"`
	Timestamp       time.Time `json:"timestamp"`
	EditedTimestamp time.Time `json:"edited_timestamp"`
	TTS             bool      `json:"tts"`
	Type            int       `json:"type"`
	WebhookID       string    `json:"webhook_id"`
}
