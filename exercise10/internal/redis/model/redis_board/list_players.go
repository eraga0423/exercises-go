package redis_board

import (
	"context"
	"log/slog"
	"strings"

	"github.com/my/repo/internal/types/pubsub"
)

func (l *LeaderBoardRedis) ListPlayers(req pubsub.ListPlayerReq) (pubsub.ListPlayerResp, error) {
	ctx := context.Background()
	topPlayers, err := l.rdb.ZRevRangeWithScores(ctx, l.conf.Redis.RedisName, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var players []pubsub.ItemPlayerResp
	for _, player := range topPlayers {
		member := player.Member.(string)
		parts := strings.SplitN(member, ":", 2)
		if len(parts) < 2 {
			slog.Error("Invalid player format")
			continue
		}
		playerID := parts[0]
		playerName := parts[1]

		players = append(players, &itemPlayersResp{
			ID:    playerID,
			Name:  playerName,
			Score: int(player.Score),
		})

	}

	return &listPlayerResp{Players: players}, nil
}

type listPlayerResp struct {
	Players []pubsub.ItemPlayerResp
}

type itemPlayersResp struct {
	ID    string
	Name  string
	Score int
}

func (p *itemPlayersResp) GetPlayerID() string {
	return p.ID
}
func (p *itemPlayersResp) GetPlayerName() string {
	return p.Name
}
func (p *itemPlayersResp) GetPlayerScore() int {
	return p.Score
}

func (p *listPlayerResp) GetList() []pubsub.ItemPlayerResp {
	return p.Players
}
