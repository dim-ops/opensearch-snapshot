package opensearch

import "go.uber.org/fx"

var Module = fx.Options(
	fx.Provide(
		NewAWSConfig,
		NewOpenSearchClient,
	),
)
