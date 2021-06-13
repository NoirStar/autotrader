package main

import (
	"github.com/noirstar/autotrader/router"
)

func main() {
	// var record *techan.TradingRecord
	// var series *techan.TimeSeries

	// candleC, err := analysis.CandleGenerator("KRW-BTC", 3, 1000)
	// utils.CheckErr(err)

	// series, record = analysis.RunDynamicStrategy(analysis.CreateDoubleBollingerStrategy, candleC)
	// totalProfit := techan.TotalProfitAnalysis{}.Analyze(record)
	// tradeCount := len(record.Trades)

	// //fmt.Println("Candle Length : ", len(series.Candles))
	// fmt.Println("Total Profit : ", totalProfit)
	// fmt.Println("Trade Count : ", tradeCount)
	// fmt.Println("Last Candle Time : ", series.LastCandle().Period.End)

	// // for idx, candle := range series.Candles {
	// // 	fmt.Println(idx, candle.ClosePrice)
	// // }

	debug := true

	router := router.New()

	if debug {
		router.Logger.Fatal(router.Start(":1323"))
	}

}
