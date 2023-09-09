package presenter

type UpdatePasswordRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u UpdatePasswordRequest) IsValidUpdatePasswordRequest() {
	// TODO - implement me
}
