package dto

// DeleteUserReq represents the request for deleting a user
type DeleteUserReq struct {
	UID string `json:"uid"`
}

// DeleteUserRes represents the response for deleting a user
type DeleteUserRes struct {
	Success bool `json:"success"`
}
