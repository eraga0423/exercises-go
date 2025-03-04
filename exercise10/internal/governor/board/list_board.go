package board

import (
	"context"
	"log/slog"
	"strings"

	"github.com/my/repo/internal/types/controller"
)

func (r *LeaderBoard) ListPlayers(req controller.ListPlayerReq) (controller.ListPlayerResp, error) {
	ctx := context.Background()
	topPlayers, err := r.rdb.ZRevRangeWithScores(ctx, r.conf.Redis.RedisName, 0, -1).Result()
	if err != nil {
		return nil, err
	}
	var players []controller.ItemPlayerResp
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
	return &ListPlayerResp{Players: players}, nil
}

type ListPlayerResp struct {
	Players []controller.ItemPlayerResp
}

func (i *ListPlayerResp) GetList() []controller.ItemPlayerResp {
	return i.Players
}

type itemPlayersResp struct {
	ID    string
	Name  string
	Score int
}

func (p *itemPlayersResp) GetID() string {
	return p.ID
}
func (p *itemPlayersResp) GetName() string {
	return p.Name
}
func (p *itemPlayersResp) GetScore() int {
	return p.Score
}
