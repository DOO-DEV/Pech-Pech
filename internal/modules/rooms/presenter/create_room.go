package presenter

type CreateRoomRequest struct {
	Description string `json:"description"`
	Category    string `json:"category"`
	Name        string `json:"name"`
}
