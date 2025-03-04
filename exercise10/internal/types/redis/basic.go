package pubsub

type LeaderBoard interface {
	CreatePlayer(CreatePlayerReq) (CreatePlayerResp, error)
	ListPlayers(ListPlayerReq) (ListPlayerResp, error)
}

type MyRedis interface {
	LeaderBoard
}
