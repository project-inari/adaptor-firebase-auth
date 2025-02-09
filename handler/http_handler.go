package handler

import (
	"context"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/project-inari/adaptor-firebase-auth/dto"
	"github.com/project-inari/adaptor-firebase-auth/pkg/request"
	"github.com/project-inari/adaptor-firebase-auth/pkg/response"
)

type httpHandler struct {
	d Dependencies
}

func newHTTPHandler(d Dependencies) *httpHandler {
	return &httpHandler{
		d: d,
	}
}

// SignUp handles the request to sign up and return a token
func (h *httpHandler) SignUp(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(dto.SignUpReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [SignUp] bad request: %v", err), "")
	}

	header := dto.SignUpReqHeader{
		AcceptLocale: c.Request().Header.Get("Accept-Locale"),
	}

	res, err := h.d.Service.SignUp(ctx, *req, header)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [SignUp] unable to sign up: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

// VerifyToken handles the request to verify a token
func (h *httpHandler) VerifyToken(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(dto.VerifyTokenReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [VerifyToken] bad request: %v", err), "")
	}

	res := h.d.Service.VerifyToken(ctx, *req)
	if !res.Success {
		return response.SuccessResponse(c, http.StatusUnauthorized, res)
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

// UpdateUsername handles the request to update a username
func (h *httpHandler) UpdateUsername(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(dto.UpdateUsernameReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [UpdateUsername] bad request: %v", err), "")
	}

	res, err := h.d.Service.UpdateUsername(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [UpdateUsername] unable to update username: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}

// DeleteUser handles the request to delete a user
func (h *httpHandler) DeleteUser(c echo.Context) error {
	ctx := context.Background()
	wrapper := request.ContextWrapper(c)

	req := new(dto.DeleteUserReq)
	if err := wrapper.Bind(req); err != nil {
		return response.ErrorResponse(c, http.StatusBadRequest, fmt.Sprintf("error - [DeleteUser] bad request: %v", err), "")
	}

	res, err := h.d.Service.DeleteUser(ctx, *req)
	if err != nil {
		return response.ErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("error - [DeleteUser] unable to delete user: %v", err), "")
	}

	return response.SuccessResponse(c, http.StatusOK, res)
}
