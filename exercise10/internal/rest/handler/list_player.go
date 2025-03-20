package handler

import (
	"fmt"
	"net/http"
)

func (h *Handler) ListPlayers(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ListPlayers")

}
