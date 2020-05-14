package commands

import (
	"errors"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/coolboy/go-robinhood"
	"github.com/weihesdlegend/Mew/clients"
	"github.com/weihesdlegend/Mew/utils"
)

const (
	TickerSeparator = "_"
)

// generate order summary for user to confirm
func previewHelper(ticker string, transactionType robinhood.OrderType, side robinhood.OrderSide, quantity uint64, price float64) (err error) {
	// to simplify testing
	if reflect.ValueOf(BufferReader).IsNil() {
		return nil
	}

	fmt.Printf("Please confirm the order details.\n"+
		"Order type: %s %s\t"+
		"Security: %s\t"+
		"Quantity: %d\t"+
		"price: %.2f [y/n]",
		transactionType, side, ticker, quantity, price)

	// wait for user confirmation
	for {
		text, _ := BufferReader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		text = strings.ToLower(text)
		if text == "y" {
			break
		} else {
			return errors.New("order is cancelled")
		}
	}
	return nil
}

// parse raw ticker string from user input
func ParseTicker(ticker string) ([]string, error) {
	ticker = strings.ToUpper(ticker)
	tickers := make([]string, 0)
	regex, _ := regexp.Compile(`[A-Z]+`)
	if strings.HasPrefix(ticker, "@") {
		fields := strings.Split(ticker, "@")
		var batch string
		if len(fields) > 1 {
			batch = fields[1]
		}
		for _, val := range strings.Split(batch, TickerSeparator) {
			if regex.MatchString(val) {
				tickers = append(tickers, val)
			}
		}
	} else {
		tickers = append(tickers, ticker)
	}
	if len(tickers) == 0 {
		return tickers, errors.New("ticker parsing error")
	}
	return tickers, nil
}

// make http calls to RH to get instrument data and current security pricing
// generate order options
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