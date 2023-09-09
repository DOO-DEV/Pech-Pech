package presenter

type ResetPasswordRequest struct {
	Code            string `json:"code"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}
