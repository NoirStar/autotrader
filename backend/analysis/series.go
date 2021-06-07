package analysis

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/noirstar/autotrading/backend/apis/restapi"
	"github.com/noirstar/autotrading/backend/models"
	"github.com/sdcoffey/big"
	"github.com/sdcoffey/techan"
)

func GetCandleData(market string, minute int, count int) (candleData []*models.ResMinuteCandles, err error) {
	data := make([]*models.ResMinuteCandles, 0)
	if count <= 200 {
		req := models.ReqMinuteCandles{
			Market: market,
			Count:  count,
		}
		reqText := restapi.GetMinuteCandles(&req, minute)
		if err := json.Unmarshal(reqText, &data); err != nil {
			return nil, err
		}
		return data, nil
	}
}

// CandleGenerator makes candles
func CandleGenerator(market string, minute int, count int) (candleC chan *techan.Candle, err error) {
	candleC = make(chan *techan.Candle)

	req := models.ReqMinuteCandles{
		Market: market,
		Count:  count,
	}
	data := make([]*models.ResMinuteCandles, 0)

	reqText := restapi.GetMinuteCandles(&req, minute)
	if err := json.Unmarshal(reqText, &data); err != nil {
		return nil, err
	}

	go func() {
		defer close(candleC)

		layout := strings.Split(time.RFC3339, "Z")[0]

		for i := len(data) - 1; i >= 0; i-- {
			start, _ := time.Parse(layout, data[i].CandleDateTimeKST)
			candle := &techan.Candle{
				Period:     techan.NewTimePeriod(start.Add(-time.Minute), time.Minute),
				OpenPrice:  big.NewFromString(fmt.Sprintf("%v", data[i].OpeningPrice)),
				ClosePrice: big.NewFromString(fmt.Sprintf("%v", data[i].TradePrice)),
				MaxPrice:   big.NewFromString(fmt.Sprintf("%v", data[i].HighPrice)),
				MinPrice:   big.NewFromString(fmt.Sprintf("%v", data[i].LowPrice)),
			}
			candleC <- candle
		}
	}()
	return candleC, nil
}

// RunDynamicStrategy run strategy
func RunDynamicStrategy(f DynamicStrategyFunc, candleC chan *techan.Candle) (*techan.TimeSeries, *techan.TradingRecord) {
	series := techan.NewTimeSeries()
	record := techan.NewTradingRecord()
	uuid := uuid.NewString()
	strategy := f(series)

	for candle := range candleC {
		series.AddCandle(candle)
		if strategy.ShouldEnter(series.LastIndex(), record) {
			fmt.Println("Enter price:", candle.ClosePrice.FormattedString(0))
			fmt.Println("Enter Time:", candle.Period.End.Format(time.Kitchen))
			record.Operate(techan.Order{
				Side:          techan.BUY,
				Security:      uuid,
				Price:         candle.ClosePrice,
				Amount:        big.NewDecimal(0.1),
				ExecutionTime: candle.Period.End,
			})
		} else if record.CurrentPosition().IsLong() && strategy.ShouldExit(series.LastIndex(), record) {
			fmt.Println("Exit price:", candle.ClosePrice.FormattedString(0))
			fmt.Println("Exit Time:", candle.Period.End.Format(time.Kitchen))
			record.Operate(techan.Order{
				Side:          techan.SELL,
				Security:      uuid,
				Price:         candle.ClosePrice,
				Amount:        big.NewDecimal(0.1),
				ExecutionTime: candle.Period.End,
			})
		}
	}
	return series, record

}
