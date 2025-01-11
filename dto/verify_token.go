package dto

// VerifyTokenReq represents the request for verifying a token with Firebase
type VerifyTokenReq struct {
	Token string `json:"token" validate:"required"`
}

// VerifyTokenRes represents the response for verifying a token with Firebase
type VerifyTokenRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Success  bool   `json:"success"`
}
