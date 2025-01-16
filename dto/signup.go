package dto

// SignUpReq represents the request for signing up a new user
type SignUpReq struct {
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	PhoneNo  string `json:"phoneNo" validate:"required"`
}

// SignUpReqHeader represents the header for signing up a new user
type SignUpReqHeader struct {
	AcceptLocale string `json:"accept-locale"`
}

// SignUpRes represents the response for signing up a new user
type SignUpRes struct {
	Username string `json:"username"`
	UID      string `json:"uid"`
	Token    string `json:"token"`
}
