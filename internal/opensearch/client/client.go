package opensearch

import (
	"log"
	"net/http"
	"time"

	cfg "github.com/dim-ops/opensearch-snapshot/internal/config"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/opensearch-project/opensearch-go"
	requestsigner "github.com/opensearch-project/opensearch-go/v2/signer/awsv2"
)

func NewOpenSearchClient(cfg *cfg.Config, awsCfg aws.Config) (clients []*opensearch.Client, err error) {
	signer, err := requestsigner.NewSignerWithService(awsCfg, "es")
	if err != nil {
		log.Fatal(err)
	}

	for i := range cfg.Opensearch.Clusters {
		client, err := opensearch.NewClient(opensearch.Config{
			Addresses: []string{cfg.Opensearch.Clusters[i]},
			Transport: &http.Transport{
				MaxIdleConns:          10,
				IdleConnTimeout:       30 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
			},
			Signer: signer,
		})
		if err != nil {
			return nil, err
		}
		clients = append(clients, client)
	}

	return clients, nil
}
