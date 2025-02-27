package board

import (
	"fmt"
	"github.com/my/repo/internal/types/controller"
	"log/slog"
)

func (r *LeaderBoard) ListPlayers(req controller.ListPlayerReq) (controller.ListPlayerResp, error) {
	if req == nil {
		slog.Error("req is nil")
		return nil, fmt.Errorf("req is nil")
	}
	response, err := r.rdb.L

	return nil, nil
}
