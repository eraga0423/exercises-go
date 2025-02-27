package redis

type ListPlayerReq interface {
	GetPlayerID() string
}
type ListPlayerResp interface {
	GetList() []ItemPlayerResp
}

type ItemPlayerResp interface {
	GetPlayerID() string
	GetPlayerName() string
	GetPlayerScore() int
}
