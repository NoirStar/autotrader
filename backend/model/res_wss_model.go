package model

// ResTickerWSS 현재가 응답 데이터 포멧
type ResTickerWSS struct {
	Type               string  `json:"type"`                  // 타입 : ticker
	Code               string  `json:"code"`                  // 마켓코드
	OpeningPrice       float64 `json:"opening_price"`         // 시가
	HighPrice          float64 `json:"high_price"`            // 고가
	LowPrice           float64 `json:"low_price"`             // 저가
	TradePrice         float64 `json:"trade_price"`           // 현재가
	PrevClosingPrice   float64 `json:"prev_closing_price"`    // 전일 종가
	Change             string  `json:"change"`                // 전일 대비 RISE :  상승 EVEN : 보합 FALL : 하락
	ChangePrice        float64 `json:"change_price"`          // 전일대비 값(절대값)
	SignedChangePrice  float64 `json:"signed_change_price"`   // 전일 대비 값
	ChangeRate         float64 `json:"change_rate"`           // 전일대비 등락률(절대값)
	SignedChangeRate   float64 `json:"signed_change_rate"`    // 전일대비 등락률
	TradeVolume        float64 `json:"trade_volume"`          // 가장 최근 거래량
	AccTradeVolume     float64 `json:"acc_trade_volume"`      // 누적거래량(UTC 0시)
	AccTradeVolume24h  float64 `json:"acc_trade_volume_24h"`  // 누적거래량 24시
	AccTradePrice      float64 `json:"acc_trade_price"`       // 누적거래대금(UTC 0시)
	AccTradePrice24h   float64 `json:"acc_trade_price_24h"`   // 누적거래대금 24시
	TradeDate          string  `json:"trade_date"`            // yyyyMMdd 최근거래일자(UTC)
	TradeTime          string  `json:"trade_time"`            // HHmmss 최근거래시각(UTC)
	TradeTimestamp     int     `json:"trade_timestamp"`       // 체결 타임스탬프(ms)
	AskBid             string  `json:"ask_bid"`               // ASK : 매도 BID : 매수
	AccAskVolume       float64 `json:"acc_ask_volume"`        // 누적 매도량
	AccBidVolume       float64 `json:"acc_bid_volume"`        // 누적 매수량
	Highest52WeekPrice float64 `json:"highest_52_week_price"` // 52주 최고가
	Highest52WeekDate  string  `json:"highest_52_week_date"`  // 52주 최고가 달성일 yyyy-MM-dd
	Lowest52WeekPrice  float64 `json:"lowest_52_week_price"`  // 52주 최저가
	Lowest52WeekDate   string  `json:"lowest_52_week_date"`   // 52주 최저가 달성일 yyyy-MM-dd
	TradeStatus        string  `json:"trade_status"`          // 거래상태 *deprecated
	MarketState        string  `json:"market_state"`          // 거래상태 PREVIEW : 입금지원 ACTIVE : 거래지원가능 DELISTED : 거래지원종료
	MarketStateForIos  string  `json:"market_state_for_ios"`  // 거래 상태 *deprecated
	IsTradingSuspended bool    `json:"is_trading_suspended"`  // 거래 정지 여부
	DelistingDate      string  `json:"delisting_date"`        // 상장 폐지일
	MarketWarning      string  `json:"market_warning"`        // 유의 종목 여부 NONE : 해당없음 CAUTION : 투자유의
	Timestamp          int     `json:"timestamp"`             // 타임스탬프(ms)
	StreamType         string  `json:"stream_type"`           // 스트림 타입 SNAPSHOT : 스냅샷 REALTIME : 실시간
}

// ResTradeWSS 체결 응답 데이터 포멧
type ResTradeWSS struct {
	Type             string  `json:"type"`               // 타입 : trade
	Code             string  `json:"code"`               // 마켓코드
	TradePrice       float64 `json:"trade_price"`        // 체결 가격
	TradeVolume      float64 `json:"trade_volume"`       // 체결량
	AskBid           string  `json:"ask_bid"`            // 매수/매도 구분 ASK : 매도 BID : 매수
	PrevClosingPrice float64 `json:"prev_closing_price"` // 전일 종가
	Change           string  `json:"change"`             // 전일 대비 RISE :  상승 EVEN : 보합 FALL : 하락
	ChangePrice      float64 `json:"change_price"`       // 부호 없는 전일 대비 값
	TradeDate        string  `json:"trade_date"`         // 체결 일자(UTC 기준)
	TradeTime        string  `json:"trade_time"`         // 체결 시각(UTC 기준)
	TradeTimestamp   int     `json:"trade_timestamp"`    // 체결 타임스탬프 (millisecond)
	Timestamp        int     `json:"timestamp"`          // 타임스탬프 (millisecond)
	SequentialID     int     `json:"sequential_id"`      // 체결 번호 (Unique)  - 체결의 유일성 판단을 위한 근거로 쓰일 수 있습니다. 하지만 체결의 순서를 보장하지는 못합니다.
	StreamType       string  `json:"stream_type"`        // 스트림 타입  SNAPSHOT : 스냅샷 REALTIME : 실시간
}

// ResOrderWSS 체결 응답 데이터 포멧
type ResOrderWSS struct {
	Type           string             `json:"type"`            // 타입 : orderbook
	Code           string             `json:"code"`            // 마켓코드
	TotalAskSize   float64            `json:"total_ask_size"`  // 호가 매도 총 잔량
	TotalBidSize   float64            `json:"total_bid_size"`  // 호가 매수 총 잔량
	OrderbookUnits []OrderBookUnitWSS `json:"orderbook_units"` // 호가
	Timestamp      int                `json:"timestamp"`       // 타임스탬프 (millisecond)
}

// OrderBookUnitWSS 호가 struct
type OrderBookUnitWSS struct {
	AskPrice float64 `json:"ask_price"` // 매도 호가
	BidPrice float64 `json:"bid_price"` // 매수 호가
	AskSize  float64 `json:"ask_size"`  // 매도 잔량
	BidSize  float64 `json:"bid_size"`  // 매수 잔량
}
