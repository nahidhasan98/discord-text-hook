package discordtexthook

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

type Embed struct {
	Title       string  `json:"title,omitempty"`
	Description string  `json:"description,omitempty"`
	URL         string  `json:"url,omitempty"`
	Color       int     `json:"color,omitempty"`
	Footer      *Footer `json:"footer,omitempty"`
	Image       *Image  `json:"image,omitempty"`
	Thumbnail   *Image  `json:"thumbnail,omitempty"`
	Author      *Author `json:"author,omitempty"`
	Fields      []Field `json:"fields,omitempty"`
}

type Footer struct {
	Text    string `json:"text"`
	IconURL string `json:"icon_url,omitempty"`
}

type Image struct {
	URL string `json:"url"`
}

type Author struct {
	Name    string `json:"name"`
	URL     string `json:"url,omitempty"`
	IconURL string `json:"icon_url,omitempty"`
}

type Field struct {
	Name   string `json:"name"`
	Value  string `json:"value"`
	Inline bool   `json:"inline,omitempty"`
}

type WebhookPayload struct {
	Content string  `json:"content,omitempty"`
	Embeds  []Embed `json:"embeds,omitempty"`
}
