package mail

type Mail struct {
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	AttachFiles []string
}

type MailService interface {
	CreateMail(mailReq *Mail) []byte
	SendMail(mailReq *Mail) error
}
