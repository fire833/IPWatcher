package backend

type SlackNotification struct {
}

func (n *SlackNotification) Name() string {
	return "Slack"
}

func (n *SlackNotification) Send(msg *Message) error {
	return nil
}

func (n *SlackNotification) Limit() *Limit {
	return nil
}
