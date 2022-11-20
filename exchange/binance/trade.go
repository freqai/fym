package binance

import (
	"fmt"

	binance "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/dqner/fym/exchange"
	"github.com/shopspring/decimal"
)

func (c *Client) SubscribeTrade(symbol, clientId string, responseHandler exchange.TradeHandler) {

	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := binance.WsAggTradeServe(symbol, tradeHandler(responseHandler), errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func (c *Client) SubscribeFutureTrade(symbol, clientId string, responseHandler exchange.TradeHandler) {
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsAggTradeServe(symbol, futureTradeHandler(responseHandler), errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func (c *Client) UnsubscribeTrade(symbol, clientId string) {

}

func (c *Client) UnsubscribeFutureTrade(symbol, clientId string) {

}

func tradeHandler(responseHandler exchange.TradeHandler) binance.WsAggTradeHandler {
	return func(event *binance.WsAggTradeEvent) {
		/*
			type WsAggTradeEvent struct {
				Event                 string `json:"e"`
				Time                  int64  `json:"E"`
				Symbol                string `json:"s"`
				AggTradeID            int64  `json:"a"`
				Price                 string `json:"p"`
				Quantity              string `json:"q"`
				FirstBreakdownTradeID int64  `json:"f"`
				LastBreakdownTradeID  int64  `json:"l"`
				TradeTime             int64  `json:"T"`
				IsBuyerMaker          bool   `json:"m"`
				Placeholder           bool   `json:"M"` // add this field to avoid case insensitive unmarshaling
			}
		*/

		if &event.Event != nil {
			tradeId := event.AggTradeID
			price, _ := decimal.NewFromString(event.Price)
			amount, _ := decimal.NewFromString(event.Quantity)
			tradeTimestamp := event.TradeTime
			direction := "sell"
			if event.IsBuyerMaker {
				direction = "buy"
			}

			responseHandler(exchange.TradeDetail{
				Id:        tradeId,
				Price:     price,
				Amount:    amount,
				Timestamp: tradeTimestamp,
				Direction: direction,
			})

		}

	}
}

func futureTradeHandler(responseHandler exchange.TradeHandler) futures.WsAggTradeHandler {
	return func(event *futures.WsAggTradeEvent) {
		/*
			type WsAggTradeEvent struct {
				Event            string `json:"e"`
				Time             int64  `json:"E"`
				Symbol           string `json:"s"`
				AggregateTradeID int64  `json:"a"`
				Price            string `json:"p"`
				Quantity         string `json:"q"`
				FirstTradeID     int64  `json:"f"`
				LastTradeID      int64  `json:"l"`
				TradeTime        int64  `json:"T"`
				Maker            bool   `json:"m"`
			}
		*/

		if &event.Event != nil {
			tradeId := event.AggregateTradeID
			price, _ := decimal.NewFromString(event.Price)
			amount, _ := decimal.NewFromString(event.Quantity)
			tradeTimestamp := event.TradeTime
			direction := "sell"
			if event.Maker {
				direction = "buy"
			}

			responseHandler(exchange.TradeDetail{
				Id:        tradeId,
				Price:     price,
				Amount:    amount,
				Timestamp: tradeTimestamp,
				Direction: direction,
			})

		}

	}
}
