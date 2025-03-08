package board

import (
	"github.com/my/repo/internal/types/controller"
)

func (r *LeaderBoard) ListPlayers(req controller.ListPlayerReq) (controller.ListPlayerResp, error) {
r.MyRedis.ListPlayers(req.)
	return &ListPlayerResp{}, nil
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


