package cmd

import (
	"github.com/happsie/fivem-loader/internal/config"
	"github.com/urfave/cli/v2"
)

func Setup() cli.ActionFunc {
	return func(ctx *cli.Context) error {
		path := ctx.String("server-data-path")
		return config.CreateConfig(path)
	}
}
