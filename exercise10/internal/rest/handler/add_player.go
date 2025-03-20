package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func (h *Handler) NewPlayer(w http.ResponseWriter, r *http.Request) {
	//ctx := r.Context()
	newPlayer := new(AddPlayer)

	err := json.NewDecoder(r.Body).Decode(newPlayer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	fmt.Println("NewPlayer")

}

type AddPlayer struct {
	Data *AddPlayerReq `json:"data"`
}

type AddPlayerReq struct {
	ID    int    `json:"id"`
	Score int    `json:"score"`
	Name  string `json:"name"`
}
