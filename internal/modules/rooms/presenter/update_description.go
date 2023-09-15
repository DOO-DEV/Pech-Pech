package presenter

type UpdateRoomInfoRequest struct {
	OldName     string `json:"old_name"`
	NewName     string `json:"new_name"`
	Description string `json:"description"`
}

type UpdateRoomInfoResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}
