# Discord Webhook (minimal text only version)

**discord-text-hook** is a package for sending text message to discord channel via discord webhook. This is very minimal service which can **send** text message, **edit** a sent message and **delete** a sent message.


## Installation

```bash
go get github.com/nahidhasan98/discord-text-hook
```


## Usage

Here is a minimal example usage that will **send**, **edit** and **delete** a message to discord channel.

```go
package main

import (
	"fmt"
	"time"

	discordtexthook "github.com/nahidhasan98/discord-text-hook"
)

func main() {
	webhookID := "yourWebhookID"
	webhookToken := "yourWebhookToken"

	// initializing new discord-text-hook service
	discord := discordtexthook.NewDiscordTextHookService(webhookID, webhookToken)

	// sending a new message
	response, err := discord.SendMessage("Hello one")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message sent. Message ID:", response.ID)

	// sleeping for 1 second
	time.Sleep(1 * time.Second)

	// editing a sent message
	response, err = discord.EditMessage("Hello two", response.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message edited")

	// sleeping for 1 second
	time.Sleep(1 * time.Second)

	// deleting a sent message
	err = discord.DeleteMessage(response.ID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Message deleted")
}
```