package opensearch

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	cfg "github.com/dim-ops/opensearch-snapshot/internal/config"
)

func NewAWSConfig(cfg *cfg.Config) (aws.Config, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(cfg.Opensearch.Region))
	if err != nil {
		return aws.Config{}, err
	}

	stsClient := sts.NewFromConfig(awsCfg)

	creds := stscreds.NewAssumeRoleProvider(stsClient, cfg.Opensearch.RoleARN)
	awsCfg.Credentials = aws.NewCredentialsCache(creds)

	return awsCfg, nil
}
