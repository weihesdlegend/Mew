package commands

import (
	"astuart.co/go-robinhood"
	"errors"
	"github.com/weihesdlegend/Mew/clients"
	"github.com/weihesdlegend/Mew/utils"
	"strings"
)

func ParseTicker(ticker string) []string {
	ticker = strings.ToUpper(ticker)
	tickers := make([]string, 0)
	if strings.HasPrefix(ticker, "@") {
		fields := strings.Split(ticker, "@")
		var batch string
		if len(fields) > 1 {
			batch = fields[1]
		}
		for _, val := range strings.Split(batch, "_") {
			tickers = append(tickers, val)
		}
	} else {
		tickers = append(tickers, ticker)
	}
	return tickers
}

func PrepareInsAndOpts(ticker string, AmountLimit float64, PercentLimit float64, rhClient clients.Client) (Ins *robinhood.Instrument, Opts robinhood.OrderOpts, err error) {
	Ins, insErr := rhClient.GetInstrument(ticker)
	if err = insErr; err != nil {
		return
	}

	quotes, quoteErr := rhClient.GetQuote(ticker)
	if err = quoteErr; err != nil {
		return
	}
	if len(quotes) == 0 {
		err = errors.New("no quote obtained from provided security name, please check")
		return
	}
	price := quotes[0].Price()
	price, roundErr := utils.Round(price*PercentLimit/100.0, 0.01) // limit to floating point 2 digits
	if err = roundErr; err != nil {
		return
	}
	quantity := uint64(AmountLimit / price)
	Opts = robinhood.OrderOpts{
		Quantity:      quantity,
		Price:         price,
		ExtendedHours: true,          // default to allow after hour
		TimeInForce:   robinhood.GFD, // default to GoodForDay
	}

	return
}
