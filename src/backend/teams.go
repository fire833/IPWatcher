package backend

type TeamsNotification struct {
}

func (n *TeamsNotification) Name() string {
	return "Teams"
}

func (n *TeamsNotification) Send(msg *Message) error {
	return nil
}

func (n *TeamsNotification) Limit() *Limit {
	return nil
}
