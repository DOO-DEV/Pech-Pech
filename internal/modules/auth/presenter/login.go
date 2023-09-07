package presenter

import "strings"

type LoginRequest struct {
	UsernameOrEmail string `json:"username_or_email"`
	Password        string `json:"password"`
}

func (l LoginRequest) IsLoginReqValid() {
	atIdx := strings.IndexByte(l.UsernameOrEmail, '@')
	if atIdx == 0 {
		return
	}

	return
}

type LoginResponse struct {
	UserId   string `json:"userId"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
