package handler

import (
	"context"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dim-ops/opensearch-snapshot/internal/config"
	"github.com/dim-ops/opensearch-snapshot/internal/opensearch/snapshot"
	"github.com/opensearch-project/opensearch-go/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewHandler(clients []*opensearch.Client, cfg *config.Config, log *zap.Logger) lambda.Handler {
	return lambda.NewHandler(func(ctx context.Context) (string, error) {

		for i := range clients {
			err := snapshot.CreateRepository(i, clients[i], cfg)
			if err != nil {
				log.Error("Impossible to create repository snapshot", zap.Error(err))
				return "KO", err
			}

			log.Info("Opensearch repository is created")

			err = snapshot.CreateSnapshot(clients[i])
			if err != nil {
				log.Error("Impossible to create snapshot", zap.Error(err))
				return "KO", err
			}

			log.Info("Snapshot is triggered")
		}

		return "OK", nil
	})
}

func RegisterLambdaHandler(lc fx.Lifecycle, s fx.Shutdowner, handler lambda.Handler) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go lambda.StartWithOptions(
				handler,
				lambda.WithEnableSIGTERM(func() {
					log.Println("shutdown from lambda handler")
					s.Shutdown()
				}),
			)
			return nil
		},
	})
}
