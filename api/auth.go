package api

import "github.com/gin-gonic/gin"

type loginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type loginUserResponse struct {
	AccessToken string `json:"access_token"`
	Username    string `json:"username"`
}

func Login(ctx *gin.Context) {}
