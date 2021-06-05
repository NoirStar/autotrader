package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/noirstar/autotrading/backend/apis/restapi"
	"github.com/noirstar/autotrading/backend/models"
	"github.com/noirstar/autotrading/backend/utils/myerr"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

func main() {

	req := models.ReqMinuteCandles{
		Market: "KRW-BTC",
		Count:  200,
	}
	res := []models.ResMinuteCandles{}

	reqText := restapi.GetMinuteCandles(&req, 1)
	json.Unmarshal(reqText, &res)

	series := MakeSeries(res)

	closePrices := techan.NewClosePriceIndicator(series)
	movingAverage := techan.NewSimpleMovingAverage(closePrices, 10)

	record := techan.NewTradingRecord()
	// record.Operate(order)

	entryRule := techan.And(
		// 종가가 ma를 넘어서거나, 새 포지션일때
		techan.NewCrossUpIndicatorRule(movingAverage, closePrices),
		techan.PositionNewRule{})

	exitRule := techan.And(
		// 종가가 ma 밑에 오게될때,
		techan.NewCrossDownIndicatorRule(closePrices, movingAverage),
		techan.PositionOpenRule{})

	strategy := techan.RuleStrategy{
		UnstablePeriod: 5,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}
	account := techan.BuyAndHoldAnalysis{
		TimeSeries:    series,
		StartingMoney: 1000000,
	}

	for i := 0; i < len(res); i++ {
		if strategy.ShouldEnter(i, record) {
			fmt.Println("Enter : ", i, "MA : ", movingAverage.Calculate(i).FormattedString(0), "Price : ", closePrices.Calculate(i).FormattedString(0))
			order := techan.Order{
				Side:          techan.BUY,
				Security:      uuid.NewString(),
				Price:         closePrices.Calculate(i),
				Amount:        big.NewDecimal(0.1),
				ExecutionTime: time.Now(),
			}
			record.Operate(order)
		} else if strategy.ShouldExit(i, record) {
			fmt.Println("Exit : ", i, "MA : ", movingAverage.Calculate(i).FormattedString(0), "Price : ", closePrices.Calculate(i).FormattedString(0))
			order := techan.Order{
				Side:          techan.SELL,
				Security:      uuid.NewString(),
				Price:         closePrices.Calculate(i),
				Amount:        big.NewDecimal(0.1),
				ExecutionTime: time.Now(),
			}
			record.Operate(order)
		}

	}
	fmt.Println(account.Analyze(record))

}

// MakeSeries makes new series (list of candles)
func MakeSeries(candles []models.ResMinuteCandles) *techan.TimeSeries {

	series := techan.NewTimeSeries()
	for i := len(candles) - 1; i >= 0; i-- {
		layout := strings.Split(time.RFC3339, "Z")[0]
		start, err := time.Parse(layout, candles[i].CandleDateTimeKST)
		myerr.CheckErr(err)
		period := techan.NewTimePeriod(start, time.Minute)

		candle := techan.NewCandle(period)
		candle.OpenPrice = big.NewFromString(fmt.Sprintf("%v", candles[i].OpeningPrice))
		candle.ClosePrice = big.NewFromString(fmt.Sprintf("%v", candles[i].TradePrice))
		candle.MaxPrice = big.NewFromString(fmt.Sprintf("%v", candles[i].HighPrice))
		candle.MinPrice = big.NewFromString(fmt.Sprintf("%v", candles[i].LowPrice))

		series.AddCandle(candle)
	}
	return series
}
