package infrastructure

import "context"

var MapHealthCheck = map[string]func(ctx context.Context) error{}
