package main

import (
	"fmt"

	"github.com/noirstar/autotrading/backend/analysis"
	"github.com/noirstar/autotrading/backend/utils"
	"github.com/sdcoffey/techan"
)

func main() {
	var record *techan.TradingRecord
	var series *techan.TimeSeries

	candleC, err := analysis.CandleGenerator("KRW-BTC", 1, 200)
	utils.CheckErr(err)

	series, record = analysis.RunDynamicStrategy(analysis.CreateSimpleRSIStrategy, candleC)
	totalProfit := techan.TotalProfitAnalysis{}.Analyze(record)
	tradeCount := len(record.Trades)

	fmt.Println("Total Profit : ", totalProfit)
	fmt.Println("Trade Count : ", tradeCount)
	fmt.Println("Last Candle Time : ", series.LastCandle().Period.End)

}
