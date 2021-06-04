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
	fmt.Println(res[0].TradePrice)
	series := techan.NewTimeSeries()
	for i := len(res) - 1; i >= 0; i-- {
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

	//closePrices := techan.NewClosePriceIndicator(series)
	//movingAverage := techan.NewSimpleMovingAverage(closePrices, 10)

	indicator := techan.NewClosePriceIndicator(series)

	order := techan.Order{
		Side:          techan.BUY,
		Security:      "what",
		Price:         big.NewDecimal(res[0].TradePrice),
		Amount:        big.NewDecimal(1),
		ExecutionTime: time.Now(),
	}

	record := techan.NewTradingRecord()
	record.Operate(order)

	entryConstant := techan.NewConstantIndicator(30)
	exitConstant := techan.NewConstantIndicator(10)

	entryRule := techan.And(
		techan.NewCrossUpIndicatorRule(entryConstant, indicator),
		techan.PositionNewRule{})

	exitRule := techan.And(
		techan.NewCrossDownIndicatorRule(indicator, exitConstant),
		techan.PositionOpenRule{})

	strategy := techan.RuleStrategy{
		UnstablePeriod: 10,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}

	for i := 0; i < len(res); i++ {
		if strategy.ShouldEnter(i, record) {
			fmt.Println(i, strategy.ShouldEnter(i, record))
		}
	}

}
