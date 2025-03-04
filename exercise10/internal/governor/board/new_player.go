package board

import "github.com/my/repo/internal/types/controller"

func (l *LeaderBoard) NewPlayer(req controller.NewPlayerReq) (controller.NewPlayerResp, error) {
	newPlayer := CreateNewPlayer(req.GetName(), req.GetId(), req.GetScore())
	l.MyRedis.CreatePlayer(newPlayer)

	return nil, nil
}

type newPlayerReq struct {
	controller.NewPlayerReq
	id    string
	name  string
	score int
}

func CreateNewPlayer(name string, id string, score int) *newPlayerReq {
	return &newPlayerReq{
		id:    id,
		name:  name,
		score: score,
	}
}

func (r newPlayerReq) GetName() string {
	return r.name
}

func (r newPlayerReq) GetIdd() string {
	return r.id
}
func (r newPlayerReq) GetScore() int {
	return r.score
}
