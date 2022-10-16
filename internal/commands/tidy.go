package commands

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli"

	"github.com/photoprism/photoprism/internal/config"
	"github.com/photoprism/photoprism/internal/photoprism"
	"github.com/photoprism/photoprism/internal/service"
)

var TidyCommand = cli.Command{
	Name:   "tidy",
	Usage:  "Tidy import path",
	Action: tidyAction,
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "path, p",
			Usage: "import path",
		},
	},
}

func tidyAction(ctx *cli.Context) error {
	conf := config.NewConfig(ctx)

	conf.SetLogLevel(logrus.FatalLevel)
	service.SetConfig(conf)

	if err := conf.Init(); err != nil {
		log.Debug(err)
	}

	// rows, cols := conf.Report()

	// result, err := report.RenderFormat(rows, cols, report.CliFormat(ctx))

	// fmt.Println(result)

	p := ctx.String("path")
	fmt.Println(p)
	f, err := photoprism.NewMediaFile(p)
	if err != nil {
		panic(err)
	}
	fmt.Printf("name: %s\n", f.CanonicalName())

	return err
}
