package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/noirstar/autotrading/backend/apis/restapi"
	"github.com/noirstar/autotrading/backend/models"
	"github.com/noirstar/autotrading/backend/utils/myerr"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

func main() {

	//accessKey := env.GetEnv("UPBIT_ACCESS_KEY")
	//secretKey := env.GetEnv("UPBIT_SECRET_KEY")

	req := models.ReqMinuteCandles{
		Market: "KRW-BTC",
		Count:  200,
	}
	res := []models.ResMinuteCandles{}

	reqText := restapi.GetMinuteCandles(&req, 1)
	json.Unmarshal(reqText, &res)

	fmt.Println(res[0])

	series := techan.NewTimeSeries()
	for i := 0; i < len(res); i++ {
		layout := strings.Split(time.RFC3339, "Z")[0]
		start, err := time.Parse(layout, res[i].CandleDateTimeKST)
		myerr.CheckErr(err)
		period := techan.NewTimePeriod(start, time.Minute)

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(fmt.Sprintf("%v", res[i].OpeningPrice))
		candle.ClosePrice = big.NewFromString(fmt.Sprintf("%v", res[i].TradePrice))
		candle.MaxPrice = big.NewFromString(fmt.Sprintf("%v", res[i].HighPrice))
		candle.MinPrice = big.NewFromString(fmt.Sprintf("%v", res[i].LowPrice))

		series.AddCandle(candle)
	}

	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewSimpleMovingAverage(closePrices, 10) // Create an exponential moving average with a window of 10

	//fmt.Println(movingAverage.Calculate(100).FormattedString(2))

	for i := 0; i < len(res); i++ {
		fmt.Println(movingAverage.Calculate(i).FormattedString(0))
	}
}
