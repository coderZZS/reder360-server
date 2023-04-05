package controller

import "gin-cli/src/models"

type _ResponsePostList struct {
	Code    ResCode              `json:"code"`
	Message string               `json:"message"`
	Data    []*models.PostDetail `json:"data"`
}
