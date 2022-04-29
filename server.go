package server

import (
	"context"
)

type ScoreServiceServer interface {
	ListMatches(ctx context.Context, *pb.ListMatchesRe)
}
