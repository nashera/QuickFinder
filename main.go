package main

import (
	"log"
	"os"

	"github.com/nashera/QuickFinder/action"
	"github.com/nashera/QuickFinder/cache"
	"github.com/nashera/QuickFinder/finderconfig"
	"github.com/urfave/cli/v2"
)

func main() {
	myconfig := finderconfig.GetConfig()
	app := &cli.App{
		Name:  "QuickFinder",
		Usage: "view or copy clinical report quickly",
		Commands: []*cli.Command{
			{
				Name:    "buildDB",
				Aliases: []string{"b"},
				Usage:   "build a database of the path to report",
				Flags: []cli.Flag{
					&cli.BoolFlag{Name: "update", Aliases: []string{"u"}},
				},
				Action: func(c *cli.Context) error {
					if c.Bool("update") {
						// fmt.Println("update:", c.String("query"))
						// cache.BuildLocalDB(config.ReportFolder, DBPath)
						cache.UpdateLocalDB(myconfig.ReportFolder, myconfig.DBPath)

					} else {
						// fmt.Println("initialize")
						cache.BuildLocalDB(myconfig.ReportFolder, myconfig.DBPath)
					}

					return nil
				},
			},
			{
				Name:    "copy",
				Aliases: []string{"c"},
				Usage:   "copy report to the output folder",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "query", Aliases: []string{"q"}, Required: true},
				},
				Action: func(c *cli.Context) error {
					// fmt.Println("query:", c.String("query"))
					dc, _ := cache.ConnectDB(myconfig.DBPath)
					defer dc.CloseDB()
					var results = dc.QueryResult(c.String("query"))
					action.CopyToOutput(results, myconfig.OutputFolder)
					return nil
				},
			},
			{
				Name:    "view",
				Aliases: []string{"v"},
				Usage:   "view the information of report or sample",
				Subcommands: []*cli.Command{
					{
						Name:  "report",
						Usage: "view information of report",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "query", Aliases: []string{"q"}, Required: true},
						},
						Action: func(c *cli.Context) error {
							dc, _ := cache.ConnectDB(myconfig.DBPath)
							defer dc.CloseDB()
							var results = dc.QueryResult(c.String("query"))
							action.ViewReports(results)
							return nil
						},
					},
					{
						Name:  "sample",
						Usage: "view information of sample ",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "query", Aliases: []string{"q"}, Required: true},
						},
						Action: func(c *cli.Context) error {
							dc, _ := cache.ConnectDB(myconfig.DBPath)
							defer dc.CloseDB()
							var samples = dc.QuerySample(c.String("query"))
							action.ViewSamples(samples)
							return nil
						},
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
