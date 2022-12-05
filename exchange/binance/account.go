package binance

/*
func (c *Client) SubscribeFutureAccount(symbol, clientId string, responseHandler exchange.FutureAccountHandler) {

	errHandler := func(err error) {
		fmt.Println(err)
	}
	doneC, _, err := futures.WsUserDataServe(symbol, futureAccountHandler(responseHandler), errHandler)
	if err != nil {
		fmt.Println(err)
		return
	}
	<-doneC
}

func (c *Client) UnsubscribeAccount(symbol, clientId string) {
	hb := new(marketwebsocketclient.TradeWebSocketClient).Init(c.Host)
	hb.UnSubscribe(symbol, clientId)
}

func futureAccountHandler(responseHandler exchange.FutureAccountHandler) futures.WsUserDataHandler {
	return func(event *futures.WsUserDataEvent) {

		var Currencies []exchange.Balance
		var Orders []exchange.Order

		if event.Event == "ACCOUNT_UPDATE" {
			for _, v := range event.AccountUpdate.Balances {
				Currencies = append(Currencies, exchange.Balance{
					Currency:  v.Asset,
					Available: v.Balance,
					Locked:    v.Asset})

			}

			if event.Event == "ORDER_TRADE_UPDATE" {
				return
			}
			Currencies := make(map[string]exchange.FutureAccountCurrency)
			Orders := make(map[string]exchange.FutureAccountOrder)

			responseHandler(exchange.AccountBalance{"future", Currencies, Orders})

		}
	}
}
*/
