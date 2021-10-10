package backend

import "fmt"

func LogNotifyFail(n Notifier) string {
	return fmt.Sprintf("Error/s with sending %s notification: %s", n.Name(), n.Error())
}
