package controller

type LeaderBoard interface {
	ListPlayers(ListPlayerReq) (ListPlayerResp, error)
	NewPlayer(NewPlayerReq) (NewPlayerResp, error)
}

type Controller interface {
	LeaderBoard
}
