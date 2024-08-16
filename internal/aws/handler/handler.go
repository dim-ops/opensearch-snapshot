package handler

import (
	"context"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/dim-ops/opensearch-snapshot/internal/config"
	"github.com/dim-ops/opensearch-snapshot/internal/opensearch/snapshot"
	"github.com/opensearch-project/opensearch-go"
)

func Handler(client *opensearch.Client, cfg *config.Config) (string, error) {
	err := snapshot.CreateRepository(client, cfg)
	if err != nil {
		//log.Error("Impossible to create repository snapshot", zap.Error(err))
		return "KO", err
	}

	//log.Info("Opensearch repository is created")

	err = snapshot.CreateSnapshot(client)
	if err != nil {
		//log.Error("Impossible to create snapshot", zap.Error(err))
		return "KO", err
	}

	//log.Info("Snapshot is triggered")

	return "OK", nil
}

func RegisterLambdaHandler(client *opensearch.Client, cfg *config.Config) {
	lambda.Start(func(ctx context.Context) (string, error) {
		return Handler(client, cfg)
	})
}
