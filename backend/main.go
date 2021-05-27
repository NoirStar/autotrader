package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/noirstar/autotrading/backend/apis/restapi"
	"github.com/noirstar/autotrading/backend/models"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

func main() {

	//accessKey := env.GetEnv("UPBIT_ACCESS_KEY")
	//secretKey := env.GetEnv("UPBIT_SECRET_KEY")

	req := models.ReqMinuteCandles{
		Market: "KRW-BTC",
		Count:  180,
	}
	res := []models.ResMinuteCandles{}

	reqText := restapi.GetMinuteCandles(&req, 1)
	json.Unmarshal(reqText, &res)

	series := techan.NewTimeSeries()

	for i := len(res) - 1; i >= 0; i-- {
		start := res[i].CandleDateTimeKST
		period := techan.NewTimePeriod(time.Unix(start, 0), time.Minute)

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(fmt.Sprintf("%v", res[i].OpeningPrice))
		candle.ClosePrice = big.NewFromString(fmt.Sprintf("%v", res[i].TradePrice))
		candle.MaxPrice = big.NewFromString(fmt.Sprintf("%v", res[i].HighPrice))
		candle.MinPrice = big.NewFromString(fmt.Sprintf("%v", res[i].LowPrice))

		series.AddCandle(candle)
	}

	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewSimpleMovingAverage(closePrices, 10) // Create an exponential moving average with a window of 10

	fmt.Println(movingAverage.Calculate(179).FormattedString(2))
	fmt.Printf("%.2f", res[len(res)-1].TradePrice)

}
