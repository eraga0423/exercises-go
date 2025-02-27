package handler

import (
	restboard "github.com/my/repo/internal/rest/handler/rest_board"
	"github.com/my/repo/internal/types/controller"
)

type Handler struct {
	*restboard.RestBoard
}

func NewHandler(ctrl controller.Controller) *Handler {
	return &Handler{
		RestBoard: restboard.NewRestBoard(ctrl),
	}
}
