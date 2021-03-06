package tests

import (
	"github.com/coolboy/go-robinhood"
	"github.com/stretchr/testify/mock"
)

type RHClientMock struct {
	mock.Mock
}

func (c *RHClientMock) GetQuote(ticker string) (quotes []robinhood.Quote, err error) {
	args := c.Called(ticker)
	return args.Get(0).([]robinhood.Quote), args.Error(1)
}

func (c *RHClientMock) GetInstrument(ticker string) (ins *robinhood.Instrument, err error) {
	args := c.Called(ticker)
	return args.Get(0).(*robinhood.Instrument), args.Error(1)
}

func (c *RHClientMock) Order(ins *robinhood.Instrument, opts robinhood.OrderOpts) (orderOutput *robinhood.OrderOutput, err error) {
	args := c.Called(ins, opts)
	return args.Get(0).(*robinhood.OrderOutput), args.Error(1)
}

func (c *RHClientMock) GetPositions() (positions []robinhood.Position, err error) {
	args := c.Called()
	return args.Get(0).([]robinhood.Position), args.Error(1)
}

func (c *RHClientMock) GetInstrumentByURL(url string) (ins *robinhood.Instrument, err error) {
	args := c.Called(url)
	return args.Get(0).(*robinhood.Instrument), args.Error(1)
}

var rhClientMocker = new(RHClientMock)
