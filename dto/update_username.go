package dto

// UpdateUsernameReq represents the request to update the username
type UpdateUsernameReq struct {
	UID         string `json:"uid"`
	NewUsername string `json:"newUsername"`
}

// UpdateUsernameRes represents the response of updating the username
type UpdateUsernameRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Success  bool   `json:"success"`
}
