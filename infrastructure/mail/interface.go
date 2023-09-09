package mail

type IMail interface {
	SendingMail(mail *Mail) error
}
