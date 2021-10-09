package backend

type WebhookNotification struct {
}

func (n *WebhookNotification) Name() string {
	return "Webhook"
}

func (n *WebhookNotification) Send(msg *Message) error {
	return nil
}

func (n *WebhookNotification) Limit() *Limit {
	return nil
}
