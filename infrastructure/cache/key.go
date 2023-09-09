package cache

func MailOtpKey(email string) string {
	return "cache_mail_otp#" + email
}
