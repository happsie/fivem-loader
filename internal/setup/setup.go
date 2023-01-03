package setup

import (
	"github.com/urfave/cli/v2"
)

func Setup() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		path := ctx.String("server-data-path")
		return CreateConfig(path)
	}
}
