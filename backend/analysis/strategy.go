package analysis

import "github.com/sdcoffey/techan"

// DynamicStrategyFunc function type
type DynamicStrategyFunc func(*techan.TimeSeries) *techan.RuleStrategy

// CreateMACDStrategy macd
func CreateMACDStrategy(series *techan.TimeSeries) *techan.RuleStrategy {
	closePrice := techan.NewClosePriceIndicator(series)
	macd := techan.NewMACDIndicator(closePrice, 12, 26)
	macdHist := techan.NewMACDHistogramIndicator(macd, 9)

	return &techan.RuleStrategy{
		EntryRule:      techan.NewCrossUpIndicatorRule(techan.NewConstantIndicator(0), macdHist),
		ExitRule:       techan.NewCrossDownIndicatorRule(macdHist, techan.NewConstantIndicator(0)),
		UnstablePeriod: 100,
	}
}

// CreateEMAStrategy EMA 전략
func CreateEMAStrategy(series *techan.TimeSeries) *techan.RuleStrategy {
	closePrice := techan.NewClosePriceIndicator(series)
	ema15 := techan.NewEMAIndicator(closePrice, 15)
	ema7 := techan.NewEMAIndicator(closePrice, 7)

	return &techan.RuleStrategy{
		EntryRule: techan.And(
			techan.OverIndicatorRule{First: closePrice, Second: ema7},
			techan.OverIndicatorRule{First: ema7, Second: ema15},
		),
		ExitRule:       techan.UnderIndicatorRule{First: closePrice, Second: ema15},
		UnstablePeriod: 15,
	}
}

// CreateDoubleBollingerStrategy 전략 3분봉 기준
// 볼린저 35의 하방이 17하방보다 낮아졌다가, 다시 35의 하방이 17하방보다 높아졌을때 진입
// 탈출은 17 볼린저 밴드를 터치할때?
func CreateDoubleBollingerStrategy(series *techan.TimeSeries) *techan.RuleStrategy {
	closePrice := techan.NewClosePriceIndicator(series)
	bb35Upper := techan.NewBollingerUpperBandIndicator(closePrice, 35, 2)
	bb35Lower := techan.NewBollingerLowerBandIndicator(closePrice, 35, 2)
	bb17Upper := techan.NewBollingerUpperBandIndicator(closePrice, 17, 2)
	bb17Lower := techan.NewBollingerLowerBandIndicator(closePrice, 17, 2)

	return &techan.RuleStrategy{
		EntryRule: techan.NewCrossUpIndicatorRule(bb17Lower, bb35Lower),
		// ExitRule: techan.Or(
		// 	techan.OverIndicatorRule{First: closePrice, Second: bb17Upper},
		// 	techan.Or(
		// 		techan.OverIndicatorRule{First: closePrice, Second: bb35Upper},
		// 		techan.NewCrossDownIndicatorRule(bb35Lower, bb17Lower),
		// 	),
		// ),
		ExitRule: techan.Or(
			techan.OverIndicatorRule{First: closePrice, Second: bb17Upper},
			techan.OverIndicatorRule{First: closePrice, Second: bb35Upper},
		),
		UnstablePeriod: 15,
	}

}

// CreateRSIStrategy RSI 전략
func CreateRSIStrategy(series *techan.TimeSeries) *techan.RuleStrategy {
	closePrice := techan.NewClosePriceIndicator(series)
	lowRSI := techan.NewConstantIndicator(30)
	highRSI := techan.NewConstantIndicator(70)
	rsi := techan.NewRelativeStrengthIndexIndicator(closePrice, 14)
	ma9 := techan.NewSimpleMovingAverage(closePrice, 9)

	return &techan.RuleStrategy{
		EntryRule: techan.And(
			techan.UnderIndicatorRule{First: closePrice, Second: ma9},
			techan.UnderIndicatorRule{First: rsi, Second: lowRSI},
		),
		ExitRule: techan.And(
			techan.OverIndicatorRule{First: closePrice, Second: ma9},
			techan.OverIndicatorRule{First: rsi, Second: highRSI},
		),
		UnstablePeriod: 15,
	}
}

// CreateSimpleRSIStrategy 테스트 전략
func CreateSimpleRSIStrategy(series *techan.TimeSeries) *techan.RuleStrategy {
	closePrice := techan.NewClosePriceIndicator(series)
	lowRSI := techan.NewConstantIndicator(40)
	highRSI := techan.NewConstantIndicator(50)
	rsi := techan.NewRelativeStrengthIndexIndicator(closePrice, 14)

	return &techan.RuleStrategy{
		EntryRule:      techan.UnderIndicatorRule{First: rsi, Second: lowRSI},
		ExitRule:       techan.OverIndicatorRule{First: rsi, Second: highRSI},
		UnstablePeriod: 15,
	}
}

// CreateSimpleMAStrategy 테스트 전략
func CreateSimpleMAStrategy(series *techan.TimeSeries) *techan.RuleStrategy {
	closePrice := techan.NewClosePriceIndicator(series)
	ma9 := techan.NewSimpleMovingAverage(closePrice, 9)

	return &techan.RuleStrategy{
		EntryRule:      techan.UnderIndicatorRule{First: closePrice, Second: ma9},
		ExitRule:       techan.OverIndicatorRule{First: closePrice, Second: ma9},
		UnstablePeriod: 10,
	}
}
