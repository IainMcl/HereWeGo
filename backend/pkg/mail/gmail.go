package mail

type GmailSender struct {
	name              string
	fromEmailAddress  string
	fromEmailPassword string
}

func NewGmailSender(name, fromEmailAddress, fromEmailPassword string) *GmailSender {
	return &GmailSender{name, fromEmailAddress, fromEmailPassword}
}
