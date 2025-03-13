package board

import (
	"github.com/my/repo/internal/types/controller"
)

func (r *LeaderBoard) ListPlayers(req controller.ListPlayerReq) (controller.ListPlayerResp, error) {

	response, err := r.MyRedis.ListPlayers(nil)
	if err != nil {
		return nil, err
	}
	newresp := response.GetList()

	return newresp, nil
}

type ListPlayerResp struct {
	Players []controller.ItemPlayerResp
}

func newListPlayer() *ListPlayerResp {
	return &ListPlayerResp{}
}
func (i *ListPlayerResp) GetList() []controller.ItemPlayerResp {
	return i.Players
}
