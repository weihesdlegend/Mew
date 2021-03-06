package commands

import "github.com/urfave/cli/v2"

// flag destination variables
var ticker string
var shares uint64
var limit float64
var limitSell float64
var totalValue float64
var percent float64
var percentTrailing int

// flag destination for auth
var user string
var password string
var mfa string

// supported flags
var tickerFlag cli.StringFlag         // security name
var sharesFlag cli.Uint64Flag         // number of shares to buy or sell
var limitBuyFlag cli.Float64Flag      // limit buy percentage to market price
var limitSellFlag cli.Float64Flag     // limit sell percentage to market price
var totalValueFlag cli.Float64Flag    // limit order maximum transaction amount
var percentToSellFlag cli.Float64Flag // percentage of total amount to sell
var percentTrailingFlag cli.IntFlag   // percentage to use in trailing stop orders

// flags for auth
var userFlag cli.StringFlag
var passwordFlag cli.StringFlag
var mfaFlag cli.StringFlag

func init() {
	userFlag = cli.StringFlag{
		Name:        "user",
		Aliases:     []string{"u"},
		Required:    true,
		Destination: &user,
	}

	passwordFlag = cli.StringFlag{
		Name:        "password",
		Aliases:     []string{"p"},
		Required:    true,
		Destination: &password,
	}

	mfaFlag = cli.StringFlag{
		Name:        "mfa",
		Aliases:     []string{"m"},
		Required:    false,
		Destination: &mfa,
	}

	tickerFlag = cli.StringFlag{
		Name:        "ticker",
		Aliases:     []string{"t"},
		Value:       "YANG",
		Required:    true,
		Destination: &ticker,
	}

	sharesFlag = cli.Uint64Flag{
		Name:        "shares",
		Aliases:     []string{"s"},
		Value:       0,
		Required:    false,
		Destination: &shares,
	}

	limitBuyFlag = cli.Float64Flag{
		Name:        "limitforbuy",
		Aliases:     []string{"l"},
		Required:    false,
		Value:       99.0,
		Destination: &limit,
	}

	limitSellFlag = cli.Float64Flag{
		Name:        "limitforsell",
		Aliases:     []string{"ls"},
		Required:    false,
		Value:       101.0,
		Destination: &limitSell,
	}

	totalValueFlag = cli.Float64Flag{
		Name:        "value",
		Aliases:     []string{"v"},
		Required:    false,
		Value:       500.0,
		Destination: &totalValue,
	}

	percentToSellFlag = cli.Float64Flag{
		Name:        "percentToSell",
		Aliases:     []string{"ps"},
		Required:    false,
		Value:       0,
		Destination: &percent,
	}

	percentTrailingFlag = cli.IntFlag{
		Name:        "percentTrailing",
		Aliases:     []string{"pt"},
		Required:    false,
		Value:       10,
		Destination: &percentTrailing,
	}
}
