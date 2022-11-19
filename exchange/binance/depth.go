package binance

import (
	binance "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/dqner/fym/exchange"
	jsonparser "github.com/buger/jsonparser"
)

func (c *Client) SubscribeDepth(symbol, clientId string, responseHandler exchange.OrderBookHandler) {
	endpoint := fmt.Sprintf("%s/%s@bookTicker", "wss://stream.binance.com:9443/ws", strings.ToLower(symbol))
	cfg := binance.NewWsConfig(endpoint, []string{})
	return binance.WsServe(cfg, depthHandler(responseHandler), errHandler)

}

func (c *Client) SubscribeFutureDepth(symbol, clientId string, responseHandler exchange.OrderBookHandler) {
	endpoint := fmt.Sprintf("%s/%s@depth20@100ms", "wss://fstream.binance.com/ws", strings.ToLower(symbol))
	cfg := futures.NewWsConfig(endpoint, []string{})
	 
	return futures.WsServe(cfg, futureDepthHandler(responseHandler), errHandler)
}

func (c *Client) UnsubscribeDepth(symbol, clientId string) {
	hb := new(marketwebsocketclient.TradeWebSocketClient).Init(c.Host)
	hb.UnSubscribe(symbol, clientId)
}


func (c *Client) UnsubscribeFutureDepth(symbol, clientId string) {
	hb := new(marketwebsocketclient.TradeWebSocketClient).Init(c.Host)
	hb.UnSubscribe(symbol, clientId)
}



func depthHandler(orderBookHandler exchange.OrderBookHandler) futures.WsHandler {
	return func(message []byte) {
 		

		var asks []exchange.Quote
		var bids []exchange.Quote

		if &message != nil {
			 
			depthId, err := jsonparser.GetInt(message, "u")

			jsonparser.ArrayEach(message , cb func(value []byte, dataType jsonparser.ValueType, offset int, err error){
				asks = append(asks, exchange.Quote{value[0], value[1]})
			},  "a")

			jsonparser.ArrayEach(message , cb func(value []byte, dataType jsonparser.ValueType, offset int, err error){
				bids = append(bids, exchange.Quote{value[0], value[1]})
			},  "b")

			var OrderBook exchange.OrderBook
			 
			OrderBook =   exchange.OrderBook{
				Id:      depthId,
				Asks:    asks,
				Bids:    bids
			}
			
			orderBookHandler(OrderBook)
		

			 
		}
		 
	}
}

func futureDepthHandler(orderBookHandler exchange.OrderBookHandler) futures.WsHandler {
	return func(message []byte) {
 		

		var Asks []exchange.Quote
		var Bids []exchange.Quote

		if &message != nil {
			 
			idContent, valueType, offset, err := jsonparser.Get(message, "u")

			jsonparser.ArrayEach(message , cb func(value []byte, dataType jsonparser.ValueType, offset int, err error){
				Asks = append(Asks, exchange.Quote{value[0], value[1]})
			},  "a")

			jsonparser.ArrayEach(message , cb func(value []byte, dataType jsonparser.ValueType, offset int, err error){
				Bids = append(Asks, exchange.Quote{value[0], value[1]})
			},  "b")

			var OrderBook exchange.OrderBook
			 
			OrderBook =   exchange.OrderBook{
				Id:      idContent,
				Asks:    Asks,
				Bids:    Bids
			}
			
			orderBookHandler(OrderBook)
		

			 
		}
		 
	}
}
errHandler := func(err error) {

	fmt.Println(err)

}