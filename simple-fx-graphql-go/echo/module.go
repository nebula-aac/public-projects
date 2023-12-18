package echo

import "go.uber.org/fx"

var Module = fx.Options(fx.Provide(newEcho), fx.Invoke(useEcho))
