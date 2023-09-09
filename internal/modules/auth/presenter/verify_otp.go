package presenter

type VerifyResetPasswordOtpRequest struct {
	Code  string `json:"code"`
	Email string `json:"email"`
}
