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

// CreateRSIStrategy RSI 전략
func CreateRSIStrategy(series *techan.TimeSeries) *techan.RuleStrategy {
	closePrice := techan.NewClosePriceIndicator(series)
	lowRSI := techan.NewConstantIndicator(40)
	highRSI := techan.NewConstantIndicator(60)
	rsi := techan.NewRelativeStrengthIndexIndicator(closePrice, 14)
	ema7 := techan.NewEMAIndicator(closePrice, 9)

	return &techan.RuleStrategy{
		EntryRule: techan.And(
			techan.UnderIndicatorRule{First: closePrice, Second: ema7},
			techan.UnderIndicatorRule{First: rsi, Second: lowRSI},
		),
		ExitRule: techan.And(
			techan.OverIndicatorRule{First: closePrice, Second: ema7},
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

// CreateSimpleEMAStrategy 테스트 전략
func CreateSimpleEMAStrategy(series *techan.TimeSeries) *techan.RuleStrategy {
	closePrice := techan.NewClosePriceIndicator(series)
	ema7 := techan.NewEMAIndicator(closePrice, 7)

	return &techan.RuleStrategy{
		EntryRule:      techan.NewCrossUpIndicatorRule(closePrice, ema7),
		ExitRule:       techan.NewCrossDownIndicatorRule(ema7, closePrice),
		UnstablePeriod: 7,
	}
}
