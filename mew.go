package main

import (
	"os"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"github.com/weihesdlegend/Mew/commands"
)

// example input from CLI: mew buy -s 100 -t AAPL
// output: purchased 100 shares of AAPL with market order, and total cost is xxx
func main() {
	log.Info("welcome to use the Mew stock assistant")

	commands.InitCommands()

	app := cli.NewApp()
	app.Commands = []*cli.Command{
		&commands.MarketBuyCmd,
		&commands.MarketSellCmd,
		&commands.LimitBuyCmd,
		&commands.LimitSellCmd,
		&commands.AuthCmd, // Create Auth file
	}

	appRunErr := app.Run(os.Args)
	if appRunErr != nil {
		log.Fatal(appRunErr)
	}
}
