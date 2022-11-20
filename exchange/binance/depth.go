package binance

import (
	"fmt"
	"time"

	binance "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/dqner/fym/exchange"
)

func (c *Client) SubscribeDepth(symbol, clientId string, responseHandler exchange.OrderBookHandler) {

	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := binance.WsDepthServe(symbol, depthHandler(responseHandler), errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC

}

func (c *Client) SubscribeFutureDepth(symbol, clientId string, responseHandler exchange.OrderBookHandler) {
	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, stopC, err := futures.WsPartialDepthServeWithRate(symbol, 20, 100, futureDepthHandler(responseHandler), errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	// use stopC to exit
	go func() {
		time.Sleep(5 * time.Second)
		stopC <- struct{}{}
	}()
	// remove this if you do not want to be blocked here
	<-doneC
}

func (c *Client) UnsubscribeDepth(symbol, clientId string) {

}

func (c *Client) UnsubscribeFutureDepth(symbol, clientId string) {

}

func depthHandler(orderBookHandler exchange.OrderBookHandler) binance.WsDepthHandler {
	return func(event *binance.WsDepthEvent) {
		/*
			type WsDepthEvent struct {
				Event         string `json:"e"`
				Time          int64  `json:"E"`
				Symbol        string `json:"s"`
				LastUpdateID  int64  `json:"u"`
				FirstUpdateID int64  `json:"U"`
				Bids          []Bid  `json:"b"`
				Asks          []Ask  `json:"a"`
			}
		*/
		var asks []exchange.Quote
		var bids []exchange.Quote

		if event != nil {

			var OrderBook exchange.OrderBook

			for i := 0; i < len(event.Asks); i++ {
				price, quantity, _ := event.Asks[i].Parse()
				asks = append(asks, exchange.Quote{price, quantity})
			}

			for i := 0; i < len(event.Bids); i++ {
				price, quantity, _ := event.Bids[i].Parse()

				bids = append(bids, exchange.Quote{price, quantity})
			}

			OrderBook = exchange.OrderBook{
				Id:   0, //event.LastUpdateID,
				Asks: asks,
				Bids: bids,
			}

			orderBookHandler(OrderBook)

		}

	}
}

func futureDepthHandler(orderBookHandler exchange.OrderBookHandler) futures.WsDepthHandler {
	return func(event *futures.WsDepthEvent) {
		/*
			type WsDepthEvent struct {
				Event         string `json:"e"`
				Time          int64  `json:"E"`
				Symbol        string `json:"s"`
				LastUpdateID  int64  `json:"u"`
				FirstUpdateID int64  `json:"U"`
				Bids          []Bid  `json:"b"`
				Asks          []Ask  `json:"a"`
			}
		*/
		var asks []exchange.Quote
		var bids []exchange.Quote

		if event != nil {

			var OrderBook exchange.OrderBook

			for i := 0; i < len(event.Asks); i++ {
				price, quantity, _ := event.Asks[i].Parse()
				asks = append(asks, exchange.Quote{price, quantity})
			}

			for i := 0; i < len(event.Bids); i++ {
				price, quantity, _ := event.Bids[i].Parse()

				bids = append(bids, exchange.Quote{price, quantity})
			}

			OrderBook = exchange.OrderBook{
				Id:   1, //event.LastUpdateID,
				Asks: asks,
				Bids: bids,
			}

			orderBookHandler(OrderBook)

		}

	}
}
