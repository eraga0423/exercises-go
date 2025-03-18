package board

import (
	"github.com/my/repo/internal/types/controller"
)

func (r *LeaderBoard) ListPlayers(req controller.ListPlayerReq) (controller.ListPlayerResp, error) {

	response, err := r.MyRedis.ListPlayers(nil)
	if err != nil {
		return nil, err
	}
	// newresp := response.GetList()

	// return newresp, nil
	ps := make([]*Player, 0, len(response.GetList()))

	for _, it := range response.GetList() {
		ps = append(ps, newPlayer(it.GetPlayerID(), it.GetPlayerName(), it.GetPlayerScore()))
	}

	return newListPlayer(ps), nil

}

type ListPlayerResp struct {
	Players []*Player
}

func newListPlayer(ps []*Player) *ListPlayerResp {
	return &ListPlayerResp{ps}
}
func (i *ListPlayerResp) GetList() []controller.ItemPlayerResp {
	list := make([]controller.ItemPlayerResp, 0, len(i.Players))

	for _, p := range i.Players {
		list = append(list, p)
	}

	return list
}

type Player struct {
	id    string
	name  string
	score int
}

func newPlayer(id, name string, score int) *Player {
	return &Player{id, name, score}
}

func (p *Player) GetID() string {
	return p.id
}

func (p *Player) GetName() string {
	return p.name
}

func (p *Player) GetScore() int {
	return p.score
}
