package binance

/*
import (
	"fmt"
	"strconv"
	"strings"

	binance "github.com/adshao/go-binance/v2"
	"github.com/adshao/go-binance/v2/futures"
	"github.com/dqner/fym/exchange"
	"github.com/huobirdcenter/huobi_golang/logging/applogger"
	"github.com/huobirdcenter/huobi_golang/pkg/client/marketwebsocketclient"
)

func (c *Client) SubscribeAccount(symbol, clientId string, responseHandler exchange.TradeHandler) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", "wss://stream.binance.com:9443/ws", strings.ToLower(symbol))
	cfg := binance.NewWsConfig(endpoint, []string{})
	return binance.WsServe(cfg, TradeHandler(responseHandler), errHandler)
}

func (c *Client) SubscribeFutureAccount(symbol, clientId string, responseHandler exchange.TradeHandler) {
	endpoint := fmt.Sprintf("%s/%s@aggTrade", "wss://fstream.binance.com/ws", strings.ToLower(symbol))
	cfg := futures.NewWsConfig(endpoint, []string{})

	return futures.WsServe(cfg, futureTradeHandler(responseHandler), errHandler)
}

func (c *Client) UnsubscribeAccount(symbol, clientId string) {
	hb := new(marketwebsocketclient.TradeWebSocketClient).Init(c.Host)
	hb.UnSubscribe(symbol, clientId)
}

func (c *Client) UnsubscribeFutureAccount(symbol, clientId string) {
	hb := new(marketwebsocketclient.TradeWebSocketClient).Init(c.Host)
	hb.UnSubscribe(symbol, clientId)
}

func AccountHandler(responseHandler exchange.TradeHandler) binance.WsHandler {
	return func(message []byte) {

		if &message != nil {

			var details []exchange.TradeDetail
			l := len(depthResponse.Tick.Data)
			for i := l - 1; i >= 0; i-- { // 火币的交易明细是时间倒序的，新数据在前
				t := depthResponse.Tick.Data[i]
				details = append(details, exchange.TradeDetail{
					Id:        t.TradeId,
					Price:     t.Price,
					Amount:    t.Amount,
					Timestamp: t.Timestamp,
					Direction: t.Direction,
				})
			}
			responseHandler(details)

		}

	}
}

func futureAccountHandler(responseHandler exchange.TradeHandler) futures.WsHandler {
	return func(message []byte) {
		var p fastjson.Parser

		if &message != nil {
			v, err := p.Parse(message)
			if err != nil {
				//log.Fatal(err)
			}

			ask := v.GetStringBytes("a")

			bid := v.GetStringBytes("b")
			a, _ := strconv.ParseFloat(ask, 32)
			b, _ := strconv.ParseFloat(bid, 32)

			if depthResponse.Tick != nil && depthResponse.Tick.Data != nil {
				applogger.Info("WebSocket received trade update: count=%d", len(depthResponse.Tick.Data))
				var details []exchange.TradeDetail
				l := len(depthResponse.Tick.Data)
				for i := l - 1; i >= 0; i-- { // 火币的交易明细是时间倒序的，新数据在前
					t := depthResponse.Tick.Data[i]
					details = append(details, exchange.TradeDetail{
						Id:        t.TradeId,
						Price:     t.Price,
						Amount:    t.Amount,
						Timestamp: t.Timestamp,
						Direction: t.Direction,
					})
				}
				responseHandler(details)
			}

			if depthResponse.Data != nil {
				applogger.Info("WebSocket received trade data: count=%d", len(depthResponse.Data))
				//for _, t := range depthResponse.Data {
				//	applogger.Info("Trade data, id: %d, price: %v, amount: %v", t.TradeId, t.Price, t.Amount)
				//}
			}
		}

	}
}
func errHandler(err error) {

	fmt.Println(err)

}
*/
