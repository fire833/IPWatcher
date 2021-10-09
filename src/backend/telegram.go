package backend

type TelegramNotification struct {
}

func (n *TelegramNotification) Name() string {
	return "Telegram"
}

func (n *TelegramNotification) Send(msg *Message) error {
	return nil
}

func (n *TelegramNotification) Limit() *Limit {
	return nil
}
