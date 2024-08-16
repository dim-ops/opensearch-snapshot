package snapshot

import "go.uber.org/fx"

var Module = fx.Module("snapshot",
	fx.Provide(
		fx.Annotate(
			fx.As(new(Snapshot)),
		),
	),
)
