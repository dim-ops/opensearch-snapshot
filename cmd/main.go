package main

import (
	"github.com/dim-ops/opensearch-snapshot/internal/aws/handler"
	"github.com/dim-ops/opensearch-snapshot/internal/config"
	opensearch "github.com/dim-ops/opensearch-snapshot/internal/opensearch/client"
	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		config.Module,
		opensearch.Module,
		handler.Module,
		fx.Provide(
			zap.NewProduction,
		),
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Invoke(handler.RegisterLambdaHandler),
	).Run()
}
