package binance

import (
	binance "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/dqner/fym/exchange"
	jsonparser "github.com/buger/jsonparser"
)

func (c *Client) SubscribeTrade(symbol, clientId string, responseHandler exchange.TradeHandler) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", "wss://stream.binance.com:9443/ws", strings.ToLower(symbol))
	cfg := binance.NewWsConfig(endpoint, []string{})
	return binance.WsServe(cfg, TradeHandler(responseHandler), errHandler)
}

func (c *Client) SubscribeFutureTrade(symbol, clientId string, responseHandler exchange.TradeHandler) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", "wss://fstream.binance.com/ws", strings.ToLower(symbol))
	cfg := futures.NewWsConfig(endpoint, []string{})
	 
	return futures.WsServe(cfg, futureTradeHandler(responseHandler), errHandler)
}

func (c *Client) UnsubscribeTrade(symbol, clientId string) {
	hb := new(marketwebsocketclient.TradeWebSocketClient).Init(c.Host)
	hb.UnSubscribe(symbol, clientId)
}


func (c *Client) UnsubscribeFutureTrade(symbol, clientId string) {
	hb := new(marketwebsocketclient.TradeWebSocketClient).Init(c.Host)
	hb.UnSubscribe(symbol, clientId)
}

func tradeHandler(responseHandler exchange.TradeHandler) binance.WsHandler {
	return func(message []byte) {
 		
		if &message != nil {
			tradeId, err :=  jsonparser.GetInt(message, "a")
			price, err :=  jsonparser.GetFloat(message, "p")
			amount, err :=  jsonparser.GetFloat(message, "q")
			tradeTimestamp, err :=  jsonparser.GetInt(message, "T")
			direction, err :=  jsonparser.GetBoolean(message, "m")
			
			responseHandler(exchange.TradeDetail{
				Id:        tradeId,
				Price:     price,
				Amount:    amount,
				Timestamp: tradeTimestamp,
				Direction: direction ? "buy" : "sell"
			})
			 
		}
		 
	}
}


func futureTradeHandler(responseHandler exchange.TradeHandler) futures.WsHandler {
	return func(message []byte) {
		
		if &message != nil {
			tradeId, err :=  jsonparser.GetInt(message, "a")
			price, err :=  jsonparser.GetFloat(message, "p")
			amount, err :=  jsonparser.GetFloat(message, "q")
			tradeTimestamp, err :=  jsonparser.GetInt(message, "T")
			direction, err :=  jsonparser.GetBoolean(message, "m")
			
			responseHandler(exchange.TradeDetail{
				Id:        tradeId,
				Price:     price,
				Amount:    amount,
				Timestamp: tradeTimestamp,
				Direction: direction ? "buy" : "sell"
			})
			 
		}
		 
		 
	}
}
errHandler := func(err error) {

	fmt.Println(err)

}
 