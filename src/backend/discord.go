package backend

type DiscordNotification struct {
}

func (n *DiscordNotification) Name() string {
	return "Discord"
}

func (n *DiscordNotification) Send(msg *Message) error {
	return nil
}

func (n *DiscordNotification) Limit() *Limit {
	return nil
}
