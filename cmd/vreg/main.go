package main

import (
	"github.com/alexcrownus/vreg"
	"github.com/caarlos0/spin"
	"github.com/urfave/cli"
	"os"
)

var version = "master"

func main() {
	//os.Setenv("HTTP_PROXY", "http://172.16.10.20:8080")
	app := cli.NewApp()
	app.Name = "vreg"
	app.Version = version
	app.Author = "Adesegun Adeyemo (alexcrownus@gmail.com)"
	app.Usage = "Command line tool for http://www.lsmvaapvs.org/index.php for checking vehicle registration details, just for fun"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "vpn",
			Value: "JJJ895AX",
			Usage: "vehicle plate number",
		},
	}
	app.Action = func(c *cli.Context) error {
		var spin = spin.New("\033[36m %s .\033[m")
		spin.Start()
		defer spin.Stop()
		err := vreg.Query(c.String("vpn"))
		if err != nil {
			spin.Stop()
			return cli.NewExitError(err.Error(), 1)
		}
		return nil
	}
	if err := app.Run(os.Args); err != nil {
		panic(err)
	}
}
