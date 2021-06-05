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

	series := MakeSeries(res)

	closePrices := techan.NewClosePriceIndicator(series)
	//movingAverage := techan.NewEMAIndicator(closePrices, 9)
	rsi := techan.NewRelativeStrengthIndexIndicator(closePrices, 14)
	//entryConstant := techan.NewConstantIndicator(30)

	record := techan.NewTradingRecord()
	// record.Operate(order)

	rsiUpRule := techan.NewCrossDownIndicatorRule(techan.NewConstantIndicator(45), rsi)
	//emaUpRule := techan.NewCrossUpIndicatorRule(movingAverage, closePrices)

	//combineRule := techan.And(rsiUpRule, emaUpRule)

	rsiDownRule := techan.NewCrossUpIndicatorRule(techan.NewConstantIndicator(60), rsi)

	entryRule := techan.And(
		rsiUpRule,
		techan.PositionNewRule{})

	exitRule := techan.And(
		rsiDownRule,
		techan.PositionOpenRule{})

	strategy := techan.RuleStrategy{
		UnstablePeriod: 15,
		EntryRule:      entryRule,
		ExitRule:       exitRule,
	}
	account := techan.BuyAndHoldAnalysis{
		TimeSeries:    series,
		StartingMoney: 1000000,
	}

	for i := 0; i < len(res); i++ {
		if strategy.ShouldEnter(i, record) {
			fmt.Println("Enter : ", i, "rsi : ", rsi.Calculate(i).FormattedString(0), "Price : ", closePrices.Calculate(i).FormattedString(0))
			order := techan.Order{
				Side:          techan.BUY,
				Security:      uuid.NewString(),
				Price:         closePrices.Calculate(i),
				Amount:        big.NewDecimal(0.1),
				ExecutionTime: time.Now(),
			}
			record.Operate(order)

		} else if strategy.ShouldExit(i, record) {
			fmt.Println("Exit : ", i, "rsi : ", rsi.Calculate(i).FormattedString(0), "Price : ", closePrices.Calculate(i).FormattedString(0))
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

func candleGenerator(market string, minute int, count int) (candleC chan *techan.Candle, err error) {
	candleC = make(chan *techan.Candle)

	req := models.ReqMinuteCandles{
		Market: market,
		Count:  count,
	}
	data := make([]*models.ResMinuteCandles, 0)

	reqText := restapi.GetMinuteCandles(&req, minute)
	json.Unmarshal(reqText, &data)

	data = data[len(data)-count:]

}
